package reader

import (
	"bufio"
	"os"
)

type LogReader interface {
	ReadFileByLine() (chan string, chan struct{})
}

type logReader struct {
	filePath   string
	streamSize int
}

func NewLogReader(path string, size int) LogReader {
	return &logReader{
		filePath:   path,
		streamSize: size,
	}
}

func (l *logReader) ReadFileByLine() (chan string, chan struct{}) {
	dataStream := make(chan string, l.streamSize)
	closeCh := make(chan struct{})
	file, err := os.Open(l.filePath)
	if err != nil {
		panic(err)
	}
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			dataStream <- scanner.Text()
		}
		close(dataStream)
		closeCh <- struct{}{}
		close(closeCh)
	}()
	return dataStream, closeCh
}
