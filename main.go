package main

import (
	"github.com/brendontj/super-ganbanzo/core/domain/records"
	"github.com/brendontj/super-ganbanzo/core/usecase"
	"github.com/brendontj/super-ganbanzo/internal/parser"
	"github.com/brendontj/super-ganbanzo/internal/reader"
	"os"
)

func main() {
	filePath := os.Args[1]
	r := reader.NewLogReader(filePath, 10000)
	dataStreamCh, closeCh := r.ReadFileByLine()

	gameRecords := records.NewGameRecords()
	reportGeneratorUseCase := usecase.NewReportGenerator(gameRecords)
	parser.ParseFromDataStream(&gameRecords, dataStreamCh, closeCh)

	//GenerateMatchReport(gameRecords, 10)
	//GenerateScoreMatchReport(gameRecords, 10)
	//GenerateDeathCausesMatchReport(gameRecords, 10)
	//GenerateDeathCausesReport(gameRecords)
	//GenerateScoreReport(gameRecords)
	reportGeneratorUseCase.GenerateAllMatchesReport()
}
