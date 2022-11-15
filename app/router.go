package app

import (
	"github.com/gorilla/mux"
	middlewares "github.com/ramadhanalfarisi/go-codebase-kocak/middleware"
)

func (a *App) Routes() {
	mux := mux.NewRouter()
	v1 := mux.StrictSlash(true).PathPrefix("/v1").Subrouter()
	secure := v1.NewRoute().Subrouter()
	v1.Use(middlewares.ApiMiddleware)
	secure.Use(middlewares.AuthMiddleware)
}