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
	Type       string     `json:"type"`
	MySQL      MySQL      `json:"mysql"`
	PostgreSQL PostgreSQL `json:"postgresql"`
	SQLite     SQLite     `json:"sqlite"`
	Redis      Redis      `json:"redis"`
}
type SQLite struct {
	Path string `json:"path"`
}
type MySQL struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name,omitempty"`
}
type Redis struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DB       int    `json:"db,omitempty"`
}
type PostgreSQL struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name,omitempty"`
	SSLMode  string `json:"sslmode"`
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
