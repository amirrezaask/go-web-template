package db

import (
	"app/config"
	"app/transport/mysql"
	"app/transport/psql"
	"app/transport/sqlite"
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
		return &mysql.Mysql{}, nil
	case "postgres":
		return &psql.Postgres{}, nil
	case "sqlite3":
		return &sqlite.SQLite{}, nil
	default:
		return SQLProvider(nil), fmt.Errorf("%s is not supported as a database provider", dbType)
	}
}
