package app

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ramadhanalfarisi/go-codebase-kocak/routers"
)

type App struct {
	Router *mux.Router
	RouterSecure *mux.Router
	DB     *sql.DB
	Route *routers.Router
}

func(a *App) Run() {
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(a.Router))
}
