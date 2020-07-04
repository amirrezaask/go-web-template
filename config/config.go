package config

import (
	"log"

	"github.com/spf13/viper"
)

var C *viper.Viper

func Init() {
	C = viper.New()
	C.SetConfigName("appname")
	C.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
