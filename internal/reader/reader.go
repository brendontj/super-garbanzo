package reader

import (
	"bufio"
	"os"
)

type LogReader interface {
	ReadFileByLine(f *os.File) (chan string, chan struct{})
}

type logReader struct {
	streamSize int
}

func NewLogReader(size int) LogReader {
	return &logReader{
		streamSize: size,
	}
}

func (l *logReader) ReadFileByLine(file *os.File) (chan string, chan struct{}) {
	dataStream := make(chan string, l.streamSize)
	closeCh := make(chan struct{})

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
