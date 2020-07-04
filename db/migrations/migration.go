package migration

import (
	"app/config"
	db2 "app/db"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4"
)

func New() (*migrate.Migrate, error) {
	db ,err := db2.DB()
	if err != nil {
		return nil, err
	}
	drv, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("./db/migrations", config.C.GetString("database.name"), drv)
	if err != nil {
		return nil, err
	}
	return m, nil
}
