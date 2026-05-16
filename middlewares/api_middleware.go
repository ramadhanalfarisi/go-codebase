package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/constants"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
)

func ApiMiddleware(c fiber.Ctx) error {
	if c.Method() == "PUT" || c.Method() == "PATCH" || c.Method() == "DELETE" {
		id := c.Params("id")
		if id == "" {
			helpers.Error(fmt.Errorf("parameter :id have to entered"))
			response := helpers.Response{Code: http.StatusBadRequest, Status: "failed", Message: constants.InvalidInput}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
	}

	return c.Next()

}
