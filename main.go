package main

import (
	"flag"
	"fmt"
	"log"

	"xtream2strm/config"
	"xtream2strm/process"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = process.GetVOD(config)
	if err != nil {
		log.Fatal(err)
	}
	err = process.GetSeries(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Application executed successfully")
}
