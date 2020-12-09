package psql

import (
	"app/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	conn *sql.DB
}

func (p *Postgres) DB() (*sql.DB, error) {
	host := config.Config.DB.PostgreSQL.Host
	port := config.Config.DB.PostgreSQL.Port
	user := config.Config.DB.PostgreSQL.Username
	password := config.Config.DB.PostgreSQL.Password
	name := config.Config.DB.PostgreSQL.DBName
	sslmode := config.Config.DB.PostgreSQL.SSLMode
	if p.conn == nil {
		conn, err := postgresConnect(host, port, user, password, name, sslmode)
		if err != nil {
			return nil, err
		}
		p.conn = conn
		return p.conn, nil
	}
	if err := p.conn.Ping(); err != nil {
		return nil, err
	}
	return p.conn, nil
}

func postgresConnect(host string, port int, user, password, name, sslmode string) (*sql.DB, error) {
	conString := postgresConnectionString(host, port, user, password, name, sslmode)
	conn, err := sql.Open("postgres", conString)
	if err != nil {
		return nil, fmt.Errorf("Error in openning postgres connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error in pinging postgres database: %w", err)
	}
	return conn, nil
}

func postgresConnectionString(host string, port int, user, password, name, sslmode string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, name, sslmode)
}
