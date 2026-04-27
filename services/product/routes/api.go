package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/services/product/controller"
	"github.com/ramadhanalfarisi/go-codebase/services/product/repository"
	"github.com/ramadhanalfarisi/go-codebase/services/product/usecase"
)

func ProductAPIRoutes(db *sql.DB, app fiber.Router) {
	repo := repository.NewProductRepository(db)
	usecase := usecase.NewProductUsecase(repo)
	controller := controller.NewProductControllerAPI(usecase)

	app.Get("products", controller.GetProducts)
	app.Get("products/:id", controller.GetProductById)
	app.Patch("products/:id", controller.UpdatePatchProduct)
	app.Put("products/:id", controller.UpdatePutProduct)
	app.Post("products", controller.CreateProduct)
	app.Delete("products/:id", controller.DeleteProduct)
}
