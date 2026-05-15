package usecase

import (
	"context"
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

func (u *UserUsecase) UserRegister(ctx context.Context, model user_model.UserRegisterInput) error {
	dataUser, err := u.getUserByEmail(ctx, model.Email)
	if dataUser.Id > 0 {
		err := errors.New("email have been registered")
		helpers.Error(err)
		return err
	}

	hash, err := helpers.HashPassword(model.Password)
	if err != nil {
		helpers.Error(err)
		return errors.New("failed to generate hash")
	}
	model.Password = hash

	err = u.repository.InsertUser(ctx, model)
	if err != nil {
		helpers.Error(err)
		return errors.New("failed to register user")
	}
	return nil
}

func (u *UserUsecase) UserLogin(ctx context.Context, model user_model.UserLoginInput) (string, error) {
	dataUser, err := u.getUserByEmail(ctx, model.Email)
	err = u.validatePassword(helpers.ValidatePassword{
		PasswordHashed: dataUser.Password,
		PasswordInput:  model.Password,
	})
	if err != nil {
		return "", err
	}
	jwt, err := u.generateJWT(dataUser)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (u *UserUsecase) getUserByEmail(ctx context.Context, email string) (user_model.DataUser, error) {
	dataUser, err := u.repository.GetUserByEmail(ctx, email)
	return dataUser, err
}

func (u *UserUsecase) generateJWT(dataUser user_model.DataUser) (string, error) {
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

func (u *UserUsecase) validatePassword(dataPass helpers.ValidatePassword) error {
	match, err := helpers.VerifyPassword(dataPass)
	if err != nil {
		helpers.Error(err)
		return errors.New("failed to login")
	}
	if !match {
		err := errors.New("invalid email or password")
		helpers.Error(err)
		return err
	}
	return nil
}
