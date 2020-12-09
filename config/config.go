package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
)

var AppName = "myapp"
var Config = struct {
	DB DB `json:"db"`
}{}

type DB struct {
	MySQL      MySQL      `json:"mysql"`
	PostgreSQL PostgreSQL `json:"postgresql"`
}

type MySQL struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name,omitempty"`
}
type PostgreSQL struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name,omitempty"`
}

func Init(configPath string) error {
	bs, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, &Config)
	if err != nil {
		return err
	}
	err = envconfig.Process(AppName, &Config)
	if err != nil {
		return err
	}
	return nil
}
