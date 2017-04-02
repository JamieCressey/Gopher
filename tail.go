package main

import (
	"log"
	"github.com/hpcloud/tail"
)

func Tailer(file, grok string) {
	_ = grok

	t, _ := tail.TailFile(file, tail.Config{Follow: true, ReOpen: true})
	for line := range t.Lines {
		formatted := GrokParse(line.Text, grok)
		log.Print(formatted)
		BatchSend(formatted)
	}
}
