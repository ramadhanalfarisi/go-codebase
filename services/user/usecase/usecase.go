package usecase

import (
	"errors"

	"github.com/ramadhanalfarisi/go-codebase/helpers"
	user_model "github.com/ramadhanalfarisi/go-codebase/services/user/models"
	"github.com/ramadhanalfarisi/go-codebase/services/user/repository"
)

type UserUsecase struct {
	repository repository.UserRepositoryInterface
}

func NewUserUsecase(repo repository.UserRepositoryInterface) UserUsecaseInterface {
	return &UserUsecase{repository: repo}
}

func (u UserUsecase) UserRegister(model user_model.UserRegisterInput) error {
	hash, err := helpers.HashPassword(model.Password)
	if err != nil {
		helpers.Error(err)
		return errors.New("failed to generate hash")
	}
	model.Password = hash

	err = u.repository.InsertUser(model)
	if err != nil {
		helpers.Error(err)
		return errors.New("failed to register user")
	}
	return nil
}

func (u UserUsecase) UserLogin(model user_model.UserLoginInput) (string, error) {
	dataUser, err := u.repository.GetUserByEmail(model)
	match, err := helpers.VerifyPassword(dataUser.Password, model.Password)
	if err != nil {
		helpers.Error(err)
		return "", errors.New("failed to login")
	}
	if !match {
		return "", errors.New("invalid email or password")
	}
	jwt, err := helpers.GenerateJWT(helpers.UserDetail{
		Id:    dataUser.Id,
		Email: dataUser.Email,
		Roles: dataUser.Roles,
	}, nil)
	if err != nil {
		helpers.Error(err)
		return "", errors.New("failed to generate token")
	}
	return jwt, nil
}
