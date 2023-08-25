package main

import (
	"bufio"
	"fmt"
	"github.com/brendontj/super-ganbanzo/core/domain/records"
	"os"
	"sync"
	"time"
)

func main() {
	filePath := os.Args[1]

	dataStream := make(chan string, 10000)
	closeCh := make(chan struct{})
	gr := records.NewGameRecords()
	var wg sync.WaitGroup
	wg.Add(2)
	go ReadFileByLine(&wg, filePath, dataStream, closeCh)
	go ParseFromDataStream(&wg, &gr, dataStream, closeCh)
	wg.Wait()

	fmt.Println(gr)
}

func ReadFileByLine(wg *sync.WaitGroup, filePath string, data chan string, closeCh chan struct{}) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data <- scanner.Text()
	}
	close(data)
	closeCh <- struct{}{}
	close(closeCh)
}

func ParseFromDataStream(wg *sync.WaitGroup, gr *records.GameRecords, dataStream chan string, closeCh chan struct{}) {
	defer wg.Done()
	//var dataBufferInMemory []string

	for {
		select {
		case data := <-dataStream:
			gr.ParseRecord(data)

			//dataBufferInMemory = append(dataBufferInMemory, data)

		case <-closeCh:
			if len(dataStream) != 0 {
				continue
			}
			//fmt.Println(dataBufferInMemory)
			fmt.Println("Finishing program execution")
			return
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
