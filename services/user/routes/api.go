package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/services/user/controller"
	"github.com/ramadhanalfarisi/go-codebase/services/user/repository"
	"github.com/ramadhanalfarisi/go-codebase/services/user/usecase"
)

func UserRoutes(db *sql.DB, app fiber.Router) {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	
	app.Post("/register", userController.UserRegister)
	app.Post("/login", userController.UserLogin)
}