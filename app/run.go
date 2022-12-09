package app

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func (a *App) Run() {
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(port_app, handlers.CORS(headers, methods, origins)(a.MainRouter))
}