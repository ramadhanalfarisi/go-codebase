package repository

import (
	"context"

	user_model "github.com/ramadhanalfarisi/go-codebase/services/user/models"
)

type UserRepositoryInterface interface {
	InsertUser(ctx context.Context, input user_model.UserRegisterInput) error
	GetUserByEmail(ctx context.Context, email string) (user_model.DataUser, error)
}
