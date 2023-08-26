package main

import (
	"fmt"
	"github.com/brendontj/super-ganbanzo/core/domain/records"
	"github.com/brendontj/super-ganbanzo/internal/parser"
	"github.com/brendontj/super-ganbanzo/internal/reader"
	"os"
)

func main() {
	filePath := os.Args[1]
	r := reader.NewLogReader(filePath, 10000)
	dataStreamCh, closeCh := r.ReadFileByLine()

	gr := records.NewGameRecords()
	parser.ParseFromDataStream(&gr, dataStreamCh, closeCh)

	fmt.Println(GenerateReport(gr))
}

func GenerateReport(gameRecords records.GameRecords) string {
	
}

