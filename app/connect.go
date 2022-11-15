package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ramadhanalfarisi/go-codebase-kocak/config"
	"github.com/ramadhanalfarisi/go-codebase-kocak/helpers"
)

var host, uname, password, port, dbname string

func init() {

	if env := config.ENVIRONMMENT; env == "production" {
		port = config.PORT_PRODDUCTION
		host = config.HOST_PRODDUCTION
		uname = config.UNAME_PRODDUCTION
		password = config.PASS_PRODDUCTION
		dbname = config.DBNAME_PRODDUCTION
	} else if env == "development" {
		port = config.PORT_DEVELOPMENT
		host = config.HOST_DEVELOPMENT
		uname = config.UNAME_DEVELOPMENT
		password = config.PASS_DEVELOPMENT
		dbname = config.DBNAME_DEVELOPMENT
	} else {
		port = config.PORT_TESTING
		host = config.HOST_TESTING
		uname = config.UNAME_TESTING
		password = config.PASS_TESTING
		dbname = config.DBNAME_TESTING
	}
}

func (a *App) ConnectDB() {
	strCon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", uname, password, host, port, dbname)
	db, err := sql.Open("mysql", strCon)
	if err != nil {
		helpers.Error(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	a.DB = db
}