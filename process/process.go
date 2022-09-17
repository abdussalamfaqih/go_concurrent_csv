package process

import (
	"encoding/csv"
	"fmt"
	"go_concurrent_csv/payload"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

const (
	benefitFreeTransfer = 5

	benefitBalance = 25

	bonusBalance = 10
)

func CalculateAverageBalance(data *payload.DataOutput) {

	data.AverageBalanced = (data.Balanced + data.PreviousBalanced) / 2

	goroutineID := goid()
	data.Thread1 = goroutineID
}

func CalculateBenefitFreeTransfer(data *payload.DataOutput) {
	if data.Balanced < 100 || data.Balanced > 150 {
		return
	}

	data.FreeTransfer = benefitFreeTransfer

	goroutineID := goid()
	data.Thread2a = goroutineID
}

func CalculateBenefitBonusBalance(data *payload.DataOutput) {
	if data.Balanced <= 150 {
		return
	}

	data.Balanced += benefitBalance
	goroutineID := goid()
	data.Thread2b = goroutineID
}

func AddLimitedBonusBalance(data *payload.DataOutput) {
	if data.Id > 100 {
		return
	}

	data.Balanced += bonusBalance

	goroutineID := goid()
	data.Thread3 = goroutineID
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func PrepareInputData(data *[]*payload.DataInput, inputPath string) {
	f, err := os.Open(inputPath)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	if err := gocsv.UnmarshalFile(f, data); err != nil { // Load clients from file
		panic(err)
	}
}

func GenerateOutput(data *[]*payload.DataOutput, outputPath string) {
	fOut, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	defer fOut.Close()

	gocsv.SetCSVWriter(func(in io.Writer) *gocsv.SafeCSVWriter {
		w := csv.NewWriter(in)
		w.Comma = ';'
		return gocsv.NewSafeCSVWriter(w)
	})

	err = gocsv.MarshalFile(data, fOut)
	if err != nil {
		log.Fatal(err)
	}
}
