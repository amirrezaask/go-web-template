package db

import (
	"app/config"
	"database/sql"
	"errors"
)

func DB() (*sql.DB, error) {
	config.C.SetDefault("database.type", "sqlite3")
	db := config.C.GetString("database.type")
	host := config.C.GetString("database.host")
	port := config.C.GetString("database.port")
	user := config.C.GetString("database.user")
	password := config.C.GetString("database.password")
	name := config.C.GetString("database.name")
	switch (db) {
	case "mysql":
		return mysql(host, port, user, password, name)
	default:
		return nil, errors.New("No database selected")
	}
}
