package mysql

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
	host := config.Config.DB.MySQL.Host
	port := config.Config.DB.MySQL.Port
	user := config.Config.DB.MySQL.Username
	password := config.Config.DB.MySQL.Password
	name := config.Config.DB.MySQL.DBName
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

func mysqlConnectionString(host string, port int, username, password, db string) string {
	if host == "" {
		host = "localhost"
	}
	if port == 0 {
		port = 3306
	}
	if username == "" {
		username = "root"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", host, port, username, password, db)
}

func mysqlConnect(host string, port int, username, password, db string) (*sql.DB, error) {
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
