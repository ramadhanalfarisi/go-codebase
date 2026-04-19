package usecase

import (
	user_model "github.com/ramadhanalfarisi/go-codebase/services/user/models"
)

type UserUsecaseInterface interface {
	UserRegister(user_model.UserRegisterInput) error
	UserLogin(user_model.UserLoginInput) (string, error)
}
