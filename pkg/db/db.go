package db

import (
	"context"
	"database/sql"
)

type SQLDB interface {
	ExecContext(ctx context.Context, stmt string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

func NewSQLDB(driver, connection string) (SQLDB, error) {
	return sql.Open(driver, connection)
}
