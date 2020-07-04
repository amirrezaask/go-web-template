package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

var mysqlConn *sql.DB
var mysqlXConn *sqlx.DB

func mysqlConnect(host, port, username, password, db string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", mysqlConnectionString(host, port, username, password, db))
	if err != nil {
		return nil, fmt.Errorf("Error in openning mysql connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error in pinging mysql database: %w", err)
	}
	mysqlConn = conn
	return conn, nil

}

func mysqlXConnect(host, port, username, password, db string) (*sqlx.DB, error) {
	conn, err := sqlx.Open("mysql", mysqlConnectionString(host, port, username, password, db))
	if err != nil {
		return nil, fmt.Errorf("Error in openning mysql connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error in pinging mysql database: %w", err)
	}
	mysqlXConn = conn
	return conn, nil
}

func mysql(host, port, username, password, db string) (*sql.DB, error) {
	if mysqlConn != nil {
		err := mysqlConn.Ping()
		if err != nil {
			return mysqlConnect(host, port, username, password, db)
		}
		return mysqlConn, nil
	}
	return mysqlConnect(host, port, username, password, db)
}

func mysqlX(host, port, username, password, db string) (*sqlx.DB, error) {
	if mysqlXConn != nil {
		err := mysqlXConn.Ping()
		if err != nil {
			return mysqlXConnect(host, port, username, password, db)
		}
		return mysqlXConn, nil
	}
	return mysqlXConnect(host, port, username, password, db)

}
