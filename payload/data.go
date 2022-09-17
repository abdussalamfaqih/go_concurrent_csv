package payload

type DataInput struct {
	Id               int    `csv:"id"`
	Name             string `csv:"Nama"`
	Age              int    `csv:"Age"`
	Balanced         int    `csv:"Balanced"`
	PreviousBalanced int    `csv:"Previous Balanced"`
	AverageBalanced  int    `csv:"Average Balanced"`
	FreeTransfer     int    `csv:"Free Transfer"`
}

type DataOutput struct {
	Id               int    `csv:"id"`
	Name             string `csv:"Nama"`
	Age              int    `csv:"Age"`
	Balanced         int    `csv:"Balanced"`
	Thread2b         int    `csv:"No 2b Thread-No"`
	Thread3          int    `csv:"No 3 Thread-No"`
	PreviousBalanced int    `csv:"Previous Balanced"`
	AverageBalanced  int    `csv:"Average Balanced"`
	Thread1          int    `csv:"No 1 Thread-No"`
	FreeTransfer     int    `csv:"Free Transfer"`
	Thread2a         int    `csv:"No 2a Thread-No"`
}
