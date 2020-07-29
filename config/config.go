package config

import (
	"log"

	"github.com/spf13/viper"
)

type appConfig struct {
	*viper.Viper
}

func AppEnv() string {
	return C.GetString("app.env")
}

//C is Instance of golobby config in use.
var C *appConfig

//Init initializes the config package and loads config file.
func Init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("app")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	C = &appConfig{v}
}
