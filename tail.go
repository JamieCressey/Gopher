package main

import (
	"time"
	"log"
	"github.com/hpcloud/tail"
)

func getEpoch() time.Time {
	return time.Now()
}

func Tailer(file, grok string) {
	_ = grok

	t, _ := tail.TailFile(file, tail.Config{Follow: true, ReOpen: true})
	for line := range t.Lines {
		record := &GopherRecord{}
		formatted := GrokParse(line.Text, grok)

		record.Data = formatted
		record.Timestamp = getEpoch()

		log.Print(formatted)
		BatchSend(record)
	}
}
