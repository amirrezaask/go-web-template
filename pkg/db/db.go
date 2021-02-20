package db

import "database/sql"

func NewSQLDB(driver, connection string) (*sql.DB, error) {
	return sql.Open(driver, connection)
}
