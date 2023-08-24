package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	filePath := os.Args[1]

	dataStream := make(chan string, 10000)
	closeCh := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)
	go ReadFileByLine(&wg, filePath, dataStream, closeCh)
	go ParseData(&wg, dataStream, closeCh)
	wg.Wait()
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

func ParseData(wg *sync.WaitGroup, dataStream chan string, closeCh chan struct{}) {
	defer wg.Done()
	var dataBufferInMemory []string
	for {
		select {
		case data := <-dataStream:
			dataBufferInMemory = append(dataBufferInMemory, data)
		case <-closeCh:
			if len(dataStream) != 0 {
				continue
			}
			fmt.Println(dataBufferInMemory)
			fmt.Println("Finishing program execution")
			return
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}

}
