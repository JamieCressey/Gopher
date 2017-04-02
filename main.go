package main

import (
	"time"
	"flag"
	"sync"
	"log"
    	"log/syslog"
)

var configfile string
var cfg Config

type GopherRecord struct {
	Timestamp	time.Time		`json:"@timestamp"`
	Data		map[string]string	`json:"message"`
	Tags      	[]string               	`json:"tags,omitempty"`
}

func init() {
	logwriter, e := syslog.New(syslog.LOG_NOTICE, "Gopher")
	if e == nil {
		log.SetOutput(logwriter)
	}

	flag.StringVar(&configfile, "configfile", "/etc/gopher.conf", "Config file location. Default: /etc/gopher.conf")
}

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	cfg = ReadConfig(configfile)

	for _, value := range cfg.Files {
		wg.Add(1)
		go func() {
			Tailer(value.Location, value.Grok)
			wg.Done()
		}()
	}
	wg.Wait()
}
