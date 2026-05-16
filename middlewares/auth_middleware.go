package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/ramadhanalfarisi/go-codebase/constants"
	"github.com/ramadhanalfarisi/go-codebase/helpers"
)

func AuthMiddleware(c fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if authorization == "" {
		helpers.Error(fmt.Errorf("Authorization empty"))
		response := &helpers.Response{Code: http.StatusUnauthorized, Status: "failed", Message: constants.InvalidToken}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	if !strings.Contains(authorization, "Bearer") {
		helpers.Error(fmt.Errorf("Have to a Bearer token"))
		response := &helpers.Response{Code: http.StatusUnauthorized, Status: "failed", Message: constants.InvalidToken}
		return c.Status(http.StatusUnauthorized).JSON(response)
	} else {
		tokenString := strings.Replace(authorization, "Bearer ", "", -1)
		claims, err := helpers.ParseUserJWT(tokenString)
		if err != nil {
			helpers.Error(err)
			response := &helpers.Response{Code: http.StatusUnauthorized, Status: "failed", Message: constants.InvalidToken}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		c.Locals("userDetail", claims)
		return c.Next()
	}

}
