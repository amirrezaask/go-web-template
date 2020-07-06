package db

import (
	"database/sql"
)
type Provider interface {
	DB() (*sql.DB, error)
}
