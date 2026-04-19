package migrate

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/ramadhanalfarisi/go-codebase/config"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ramadhanalfarisi/go-codebase/db"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
)

func Migrate() {
	dbConnect := db.ConnectDB()
	driver, err := postgres.WithInstance(dbConnect, &postgres.Config{})
	if err != nil {
		helpers.Error(err)
	}
	path := config.MIGRATIONS_PATH
	fmt.Println(path)
	m, err := migrate.NewWithDatabaseInstance(
		path,
		config.DB_NAME, driver)
	if err != nil {
		helpers.Error(err)
	}
	if err := m.Up(); err != nil {
		helpers.Error(err)
	}
	fmt.Println("Migrations ran successfully")
}
