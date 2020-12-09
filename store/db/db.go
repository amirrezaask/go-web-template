package db

import (
	"app/adapters/mysql"
	"app/adapters/psql"
	"app/adapters/sqlite"
	"app/config"
	"database/sql"
	"fmt"
)

type SQLProvider interface {
	DB() (*sql.DB, error)
}

func NewSQLProvider() (SQLProvider, error) {

	dbType := config.Config.DB.Type

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
