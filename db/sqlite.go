package db

import (
	"app/config"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	conn *sql.DB
}

func (s *SQLite) DB() (*sql.DB, error) {
	path, err := config.C.GetString("database.path")
	if err != nil {
	    return nil, fmt.Errorf("could not create SQLite instance: %w", err)
	}
	if s.conn == nil {
		conn, err := sqliteConnect(path)
		if err != nil {
			return nil, err
		}
		s.conn = conn
		return s.conn, nil
	}
	if err := s.conn.Ping(); err != nil {
		return nil, err
	}
	return s.conn, nil
}

func sqliteConnect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}