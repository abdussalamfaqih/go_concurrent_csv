package main

import (
	"fmt"
	"sort"

	"go_concurrent_csv/payload"
	"go_concurrent_csv/process"

	"github.com/jinzhu/copier"
)

const (
	inputPath = "Before Eod.csv"

	outputPath = "After_Eod.csv"

	numWorkers = 8
)

func main() {
	// startTime := time.Now()

	// Spawns pipes to handle each process
	chOutput1, closeChan1 := make(chan *payload.DataOutput, numWorkers), make(chan int)

	chOutput2a, closeChan2a := make(chan *payload.DataOutput, numWorkers), make(chan int)

	chOutput2b, closeChan2b := make(chan *payload.DataOutput, numWorkers), make(chan int)

	chOutput3, closeChan3 := make(chan *payload.DataOutput, numWorkers), make(chan int)

	var resultChan = make(chan *payload.DataOutput)

	data := []*payload.DataInput{}
	process.PrepareInputData(&data, inputPath)

	result := []*payload.DataOutput{}
	err := copier.Copy(&result, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	go publishData(chOutput1, &result)

	// Spawn workers according to numWorkers number, in every process to handle each ctep in concurrent fashion
	// Worker is created using buffered channel to implement thread style.
	// Each worker communicates via channel pipe after finishing designated process to make sure no race condition happens.
	go checkCloseChan(numWorkers, chOutput2a, closeChan1)
	createWorkerPool(numWorkers, process.CalculateAverageBalance, chOutput1, chOutput2a, closeChan1)

	go checkCloseChan(numWorkers, chOutput2b, closeChan2a)
	createWorkerPool(numWorkers, process.CalculateBenefitFreeTransfer, chOutput2a, chOutput2b, closeChan2a)

	go checkCloseChan(numWorkers, chOutput3, closeChan2b)
	createWorkerPool(numWorkers, process.CalculateBenefitBonusBalance, chOutput2b, chOutput3, closeChan2b)

	go checkCloseChan(numWorkers, resultChan, closeChan3)
	createWorkerPool(numWorkers, process.AddLimitedBonusBalance, chOutput3, resultChan, closeChan3)

	output := []*payload.DataOutput{}
	for r := range resultChan {
		output = append(output, r)
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Id < output[j].Id
	})

	process.GenerateOutput(&output, outputPath)

	// endTime := time.Now()
	// diff := endTime.Sub(startTime)
	// fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

func publishData(ch chan *payload.DataOutput, input *[]*payload.DataOutput) {
	for _, d := range *input {
		ch <- d
	}
	close(ch)
}

func createWorkerPool(numWorkers int, inpFunc func(*payload.DataOutput), fromChan, target chan *payload.DataOutput, closeChan chan int) {
	for i := 0; i < numWorkers; i++ {
		go workerProcess(inpFunc, fromChan, target, closeChan)
	}
}

func workerProcess(inFunc func(*payload.DataOutput), fromChan, target chan *payload.DataOutput, closeChan chan int) {
	for d := range fromChan {
		inFunc(d)

		target <- d
	}

	closeChan <- 1
}

func checkCloseChan(numWorker int, chanToClose chan *payload.DataOutput, handlerChan chan int) {
	var container []int
	for v := range handlerChan {
		container = append(container, v)

		if len(container) == numWorker {
			close(chanToClose)
			close(handlerChan)
		}
	}
}
