package main

import (
	"github.com/golobby/config/v2/feeder"
	"log"
	"template/cmd"
	"template/config"
)

func main() {
	err := config.Init(&feeder.Json{Path: "config.json"})
	if err != nil {
		log.Fatalln(err)
	}
	if err = cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
