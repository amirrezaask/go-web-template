package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func mysqlConnectionString(host, port, username, password, db string) string {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}
	if username == "" {
		username = "root"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", host, port, username, password, db)
}

func Mysql(host, port, username, password, db string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", mysqlConnectionString(host, port, username, password, db))
	if err != nil {
		return nil, fmt.Errorf("Error in openning mysql connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error in pinging mysql database: %w", err)
	}
	return conn, nil
}
