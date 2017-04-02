package main

import (
	"os"
	"log"

	"github.com/BurntSushi/toml"
)

// Info from config file
type Config struct {
	FirehoseDeliveryStreamName	string
	Files     			[]File
}

type File struct {
	Location	string
	Grok		string
}

// Reads info from config file
func ReadConfig(configfile string) Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
