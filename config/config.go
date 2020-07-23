package config

import (
	"github.com/golobby/config"
	"github.com/golobby/config/feeder"
)

var C *config.Config

func Init() {
	config.New(config.Options{
		Feeder: feeder.Json{
			"config.json",
		},
		Env:    "",
	})
}
