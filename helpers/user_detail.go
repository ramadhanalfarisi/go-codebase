package helpers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserIdFromAPI(c fiber.Ctx) int {
	userDetail := c.Locals("userDetail").(jwt.MapClaims)
	userId := int(userDetail["userId"].(float64))
	return userId
}

func GetUserIdFromGraphql(c context.Context) int {
	userDetail := c.Value("userDetail").(UserDetail)
	return userDetail.Id
}

func GetUserDetail(claim jwt.MapClaims) UserDetail {
	userDetail := UserDetail{
		Id:    int(claim["userId"].(float64)),
		Email: claim["email"].(string),
		Roles: claim["roles"].(string),
	}
	return userDetail
}
