package parser

import (
	"github.com/brendontj/super-ganbanzo/core/domain/records"
	"time"
)

func ParseFromDataStream(gr *records.GameRecords, dataStream chan string, closeCh chan struct{}) {
	for {
		select {
		case data := <-dataStream:
			gr.ParseRecord(data)
		case <-closeCh:
			if len(dataStream) != 0 {
				continue
			}
			return
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
