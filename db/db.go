package db

import (
	"app/config"
	"database/sql"
	"fmt"
)

type SQLProvider interface {
	DB() (*sql.DB, error)
}

func NewSQLProvider() (SQLProvider, error) {
	config.C.SetDefault("database.type", "sqlite3")
	dbType := config.C.GetString("database.type")

	switch dbType {
	case "mysql":
		return &Mysql{}, nil
	case "postgres":
		return &Postgres{}, nil
	case "sqlite3":
		return &SQLite{}, nil
	default:
		return SQLProvider(nil), fmt.Errorf("%s is not supported as a database provider", dbType)
	}
}
