package controller

import "github.com/gofiber/fiber/v3"

type UserControllerInterface interface {
	UserRegister(c fiber.Ctx) error
	UserLogin(c fiber.Ctx) error
}