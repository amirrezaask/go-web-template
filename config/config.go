package config

import (
	"github.com/golobby/config"
	"log"
)

type appConfig struct {
	*config.Config
}
func (a *appConfig) AppEnv() string {
	env, err := a.GetString("env")
	if err != nil {
		log.Fatalln(err)
	}
	return env
}

//C is Instance of golobby config in use.
var C *appConfig

//Init initializes the config package and loads config file.
func Init(opts config.Options) {
	c, err := config.New(opts)
	if err != nil {
		log.Fatalln(err)
	}
	C = &appConfig{c}
}
