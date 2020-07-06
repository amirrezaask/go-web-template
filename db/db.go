package db

import (
	"app/config"
	"database/sql"
	"fmt"
)
type Provider interface {
	DB() (*sql.DB, error)
}

func NewProvider() (Provider, error) {
	config.C.Set("database.type", "sqlite3")
	dbType := config.C.GetString("database.type")
	switch dbType {
	case "mysql":
		return &Mysql{}, nil
	case "postgres":
		return &Postgres{}, nil
	case "sqlite3":
		return &Sqlite{}, nil
	default:
		return Provider(nil), fmt.Errorf("%s is not supported as a database provider", dbType)
	}
}