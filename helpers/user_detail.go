package helpers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserId(c fiber.Ctx) int {
	userDetail := c.Locals("userDetail").(jwt.MapClaims)
	userId := userDetail["userId"].(string)
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatal("Failed to convert userId to int:", err)
	}
	return userIdInt
}
