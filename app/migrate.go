package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ramadhanalfarisi/go-codebase-kocak/config"
	"github.com/ramadhanalfarisi/go-codebase-kocak/helpers"
)

func (a *App) Migrate() {
	driver, err := mysql.WithInstance(a.DB, &mysql.Config{})
	if err != nil {
		helpers.Error(err)
	}
	path := config.MIGRATIONS_PATH
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
