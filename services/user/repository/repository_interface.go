package repository

import user_model "github.com/ramadhanalfarisi/go-codebase/services/user/models"

type UserRepositoryInterface interface {
	InsertUser(input user_model.UserRegisterInput) error
	GetUserByEmail(email string) (user_model.DataUser, error)
}
