package user_repository

import "github.com/ramadhanalfarisi/go-codebase-kocak/models"

type UserInterface interface {
	UserRegister(models.UserRegister) error
	UserLogin(models.UserLogin) (models.DataLogin, error)
}