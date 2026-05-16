// This package is the main package that connect database and service

package api

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"
	_ "net/http/pprof"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/ramadhanalfarisi/go-codebase/config"
	"github.com/ramadhanalfarisi/go-codebase/drivers"
)

// Api struct which is the main struct that will connect DB and service
type Api struct {
	DB  *sql.DB
	App *fiber.App
}

func NewApi() *Api {
	dbConnect := drivers.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: []string{"Authorization", "Content-Type"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowOrigins: []string{"*"},
	}))
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))
	app.Use(limiter.New(limiter.Config{
		Max: 100,
	}))

	return &Api{
		DB:  dbConnect,
		App: app,
	}
}

func (a *Api) Run() {
	a.LoadRoutes()
	fmt.Println("Your application running on http://localhost:" + config.PORT_API)
	go func(){
		log.Println(http.ListenAndServe(config.PPROF_API_PORT, nil))
	}()
	a.App.Listen(config.PORT_API)
}
