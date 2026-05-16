package helpers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserId(c fiber.Ctx) int {
	userDetail := c.Locals("userDetail").(jwt.MapClaims)
	userId := userDetail["userId"].(int)
	return userId
}

func GetUserDetail(claim jwt.MapClaims) UserDetail {
	userDetail := UserDetail{
		Id:    claim["userId"].(int),
		Email: claim["email"].(string),
		Roles: claim["roles"].(string),
	}
	return userDetail
}
