package migration

import (
	"app/config"
	db2 "app/store/db"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func New() (*migrate.Migrate, error) {
	prv, err := db2.NewSQLProvider()
	if err != nil {
		return nil, err
	}
	db, err := prv.DB()
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
