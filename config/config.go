package config

import (
	"log"

	"github.com/spf13/viper"
)

var C *viper.Viper

func Init() {
	C = viper.New()
	C.AddConfigPath(".")
	C.SetConfigName("app")
	C.SetConfigType("yaml")
	err := C.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
