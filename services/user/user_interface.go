package user_service

import (
	"github.com/ramadhanalfarisi/go-codebase-kocak/models"
)

type UserInterface interface {
	UserRegister(models.UserRegister)
	UserLogin(models.UserLogin)
}