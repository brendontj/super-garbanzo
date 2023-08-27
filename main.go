package main

import (
	"flag"
	"fmt"
	"github.com/brendontj/super-ganbanzo/core/domain/records"
	"github.com/brendontj/super-ganbanzo/core/usecase"
	"github.com/brendontj/super-ganbanzo/internal/parser"
	"github.com/brendontj/super-ganbanzo/internal/reader"
	"log"
	"os"
)

const AllMatchesIdentifier = -1

var (
	filePath       string
	streamChSize   int
	reportType     string
	gameIdentifier int
)

func init() {
	flag.StringVar(&filePath, "file", "log-samples/qgames.log", "a string identifying the file path of the log file")
	flag.IntVar(&streamChSize, "streamSize", 1000, "a int indicating the size of the stream channel")
	flag.StringVar(&reportType, "reportType", "game", "a string identifying the type of the report")
	flag.IntVar(&gameIdentifier, "gameIdentifier", -1, "a int indicating the game to generate the report")
}

func main() {
	flag.Parse()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(fmt.Sprintf("unable to open file, reason: %v", err))
	}
	
	logReader := reader.NewLogReader(streamChSize)
	dataStreamCh, closeCh := logReader.ReadFileByLine(file)

	gameRecords := records.NewGameRecords()
	reportGeneratorUseCase := usecase.NewReportGenerator(gameRecords)
	parser.ParseFromDataStream(&gameRecords, dataStreamCh, closeCh)

	handleUserOptions(reportGeneratorUseCase)
}

func handleUserOptions(reportGeneratorUseCase usecase.ReportGenerator) {
	switch reportType {
	case "game":
		if gameIdentifier == AllMatchesIdentifier {
			reportGeneratorUseCase.GenerateAllMatchesReport()
			return
		}
		if err := reportGeneratorUseCase.GenerateMatchReport(gameIdentifier); err != nil {
			log.Fatalln(fmt.Sprintf("Unable to generate report, reason: %v", err))
		}
	case "score":
		if gameIdentifier == AllMatchesIdentifier {
			reportGeneratorUseCase.GenerateScoreReport()
			return
		}
		if err := reportGeneratorUseCase.GenerateScoreMatchReport(gameIdentifier); err != nil {
			log.Fatalln(fmt.Sprintf("Unable to generate report, reason: %v", err))
		}
	case "death-causes":
		if gameIdentifier == AllMatchesIdentifier {
			reportGeneratorUseCase.GenerateDeathCausesReport()
			return
		}
		if err := reportGeneratorUseCase.GenerateDeathCausesMatchReport(gameIdentifier); err != nil {
			log.Fatalln(fmt.Sprintf("Unable to generate report, reason: %v", err))
		}
	default:
		log.Fatalln("Unable to generate report, report type not implemented")
	}
}
