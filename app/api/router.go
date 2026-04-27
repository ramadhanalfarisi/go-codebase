package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/middlewares"
	productRoutes "github.com/ramadhanalfarisi/go-codebase/services/product/routes"
	userRoutes "github.com/ramadhanalfarisi/go-codebase/services/user/routes"
)

func (a *Api) LoadRoutes() {
	// make new route
	api := a.App.Group("/api")
	v1 := api.Group("/v1")
	auth := api.Group("/auth/v1")

	// connect to middleware
	v1.Use(middlewares.ApiMiddleware)
	v1.Use(middlewares.AuthMiddleware)
	auth.Use(middlewares.ApiMiddleware)

	// load all routes
	a.ListRoutes(v1, auth)
}

func (a *Api) ListRoutes(apiRouter fiber.Router, authRouter fiber.Router) {
	// load all routes
	userRoutes.UserRoutes(a.DB, authRouter)
	productRoutes.ProductAPIRoutes(a.DB, apiRouter)
}
