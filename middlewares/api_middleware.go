package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
)

func ApiMiddleware(c fiber.Ctx) error {
	if c.Method() == "PUT" || c.Method() == "PATCH" || c.Method() == "DELETE" {
		id := c.Params("id")
		if id == "" {
			response := helpers.Response{Code: 400, Status: "failed", Message: "parameter :id have to entered"}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
	}

	return c.Next()

}
