package routers

import (
	"github.com/gorilla/mux"
	"github.com/ramadhanalfarisi/go-codebase-kocak/controllers"
)

type Router struct {
	Router       *mux.Router
	RouterSecure *mux.Router
	Controller   *controllers.Controller
}
