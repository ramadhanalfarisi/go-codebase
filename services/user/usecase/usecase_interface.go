package usecase

import (
	"context"

	user_model "github.com/ramadhanalfarisi/go-codebase/services/user/models"
)

type UserUsecaseInterface interface {
	UserRegister(ctx context.Context, input user_model.UserRegisterInput) error
	UserLogin(ctx context.Context, input user_model.UserLoginInput) (string, error)
}
