package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ramadhanalfarisi/go-codebase-kocak/helpers"
)

func (a *App) ConnectDB() {
	strCon := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", uname, password, host, port, dbname)
	db, err := sql.Open("mysql", strCon)
	if err != nil {
		helpers.Error(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	a.DB = db
	a.Migrate()
}