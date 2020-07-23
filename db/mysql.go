package db

import (
	"app/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	conn *sql.DB
}

func (m *Mysql) DB() (*sql.DB, error) {
	host, err := config.C.GetString("database.host")
	if err != nil {
	    return nil, fmt.Errorf("could'nt create Mysql instance %w", err)
	}
	port, err := config.C.GetString("database.port")
	if err != nil {
		return nil, fmt.Errorf("could'nt create Mysql instance %w", err)
	}
	user, err := config.C.GetString("database.user")
	if err != nil {
		return nil, fmt.Errorf("could'nt create Mysql instance %w", err)
	}
	password, err := config.C.GetString("database.password")
	if err != nil {
		return nil, fmt.Errorf("could'nt create Mysql instance %w", err)
	}
	name, err := config.C.GetString("database.name")
	if err != nil {
		return nil, fmt.Errorf("could'nt create Mysql instance %w", err)
	}
	if m.conn == nil {
		conn, err := mysqlConnect(host, port, user, password, name)
		if err != nil {
			return nil, err
		}
		m.conn = conn
		return m.conn, nil
	}
	if err := m.conn.Ping(); err != nil {
		return nil, err
	}
	return m.conn, nil
}

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

func mysqlConnect(host, port, username, password, db string) (*sql.DB, error) {
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
