package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
)

func AuthMiddleware(c fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if authorization == "" {
		response := &helpers.Response{Code: 401, Status: "failed", Message: "Authorization header is required"}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	if !strings.Contains(authorization, "Bearer") {
		response := &helpers.Response{Code: 401, Status: "failed", Message: "Token must be Bearer type"}
		return c.Status(http.StatusUnauthorized).JSON(response)
	} else {
		tokenString := strings.Replace(authorization, "Bearer ", "", -1)
		claims, err := helpers.ParseUserJWT(tokenString)
		if err != nil {
			log.Fatal(err)
		}

		c.Locals("userDetail", claims)
		return c.Next()
	}

}
