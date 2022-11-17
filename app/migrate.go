package app

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/ramadhanalfarisi/go-codebase-kocak/config"
	"github.com/ramadhanalfarisi/go-codebase-kocak/helpers"
)

func (a *App) Migrate() {
	driver, err := mysql.WithInstance(a.DB, &mysql.Config{})
	if err != nil {
		helpers.Error(err)
	}
	path := config.MIGRATIONS_LOCAL_PATH
	m, err := migrate.NewWithDatabaseInstance(
		path,
		dbname, driver)
	if err != nil {
		helpers.Error(err)
	}
	if err := m.Up(); err != nil {
		helpers.Error(err)
	}
}