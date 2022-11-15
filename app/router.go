package app

import (
	"github.com/gorilla/mux"
	"github.com/ramadhanalfarisi/go-codebase-kocak/controllers"
	middlewares "github.com/ramadhanalfarisi/go-codebase-kocak/middleware"
	"github.com/ramadhanalfarisi/go-codebase-kocak/routers"
)

func (a *App) Routes() {
	// make new route
	mux := mux.NewRouter()
	v1 := mux.StrictSlash(true).PathPrefix("/v1").Subrouter()
	secure := v1.NewRoute().Subrouter()

	// connect to middleware
	v1.Use(middlewares.ApiMiddleware)
	secure.Use(middlewares.AuthMiddleware)

	// connect to objects
	controller := &controllers.Controller{}
	controller.DB = a.DB
	routes := &routers.Router{Router: v1, RouterSecure: secure, Controller: controller}

	// set app object
	a.Router = v1
	a.RouterSecure = secure
	a.Route = routes
}

func (a *App) ListRoutes() {
	a.Route.UserRouter()
}
