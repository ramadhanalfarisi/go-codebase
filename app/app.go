package app

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/ramadhanalfarisi/go-codebase-kocak/config"
	"github.com/ramadhanalfarisi/go-codebase-kocak/routers"
)

type App struct {
	MainRouter   *mux.Router
	Router       *mux.Router
	RouterSecure *mux.Router
	DB           *sql.DB
	Route        *routers.Router
}

var host, uname, password, dbname string
var port int

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
