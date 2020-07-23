package config

import (
	"log"
	"os"

	"github.com/golobby/config"
)

//C is Instance of golobby config in use.
var C *config.Config

var (
	//APPEnv is a variable that shows which environment application is running in. (dev, prod, testing)
	APPEnv = os.Getenv("APPENV")
)

//Init initializes the config package and loads config file.
func Init(opts config.Options) {
	c, err := config.New(opts)
	if err != nil {
		log.Fatalln(err)
	}
	C = c
}
