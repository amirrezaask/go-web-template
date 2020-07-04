package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("appname")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
