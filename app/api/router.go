package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/middlewares"
	"github.com/ramadhanalfarisi/go-codebase/services/user/routes"
)

func(r *Api) LoadRoutes() {
	// make new route
	api := r.App.Group("/api")
	v1 := api.Group("/v1")
	auth := api.Group("/auth/v1")

	// connect to middleware
	v1.Use(middlewares.ApiMiddleware)
	v1.Use(middlewares.AuthMiddleware)
	auth.Use(middlewares.ApiMiddleware)

	// load all routes
	r.ListRoutes(v1, auth)
}

func (a *Api) ListRoutes(apiRouter fiber.Router, authRouter fiber.Router) {
	// load all routes
	routes.UserRoutes(a.DB, authRouter)
}