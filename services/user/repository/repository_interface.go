package repository

import user_model "github.com/ramadhanalfarisi/go-codebase/services/user/models"

type UserRepositoryInterface interface {
	InsertUser(user_model.UserRegisterInput) error
	GetUserByEmail(user_model.UserLoginInput) (user_model.DataUser, error)
}
