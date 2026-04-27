package drivers

import (
	"database/sql"
	"time"

	"github.com/ramadhanalfarisi/go-codebase/config"
	"github.com/ramadhanalfarisi/go-codebase/helpers"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", config.DB_URL)
	if err != nil {
		helpers.Error(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
