package repository

import (
	"database/sql"
	"time"

	"github.com/ramadhanalfarisi/go-codebase/helpers"
	"github.com/ramadhanalfarisi/go-codebase/helpers/query_builder"
	user_model "github.com/ramadhanalfarisi/go-codebase/services/user/models"
)

type UserRepository struct {
	db          *sql.DB
	queryHelper helpers.QueryHelperInterface
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db: db, queryHelper: helpers.NewQueryHelper(db)}
}

func (u *UserRepository) InsertUser(model user_model.UserRegisterInput) error {
	query, args := query_builder.New("users").Insert("email", "roles", "password", "created_at").Values(model.Email, model.Roles, model.Password, time.Now()).Build()
	err := u.queryHelper.Insert(query, args)
	return err
}

func (u *UserRepository) GetUserByEmail(email string) (user_model.DataUser, error) {
	var dataLogin user_model.DataUser
	query, args := query_builder.New("users").Select("id", "email", "password", "roles").Where("email = ?", email).Build()
	err := u.queryHelper.Select(query, args, &dataLogin.Id, &dataLogin.Email, &dataLogin.Password, &dataLogin.Roles)
	if err != nil {
		helpers.Error(err)
		return user_model.DataUser{}, err
	} else {
		return dataLogin, nil
	}
}
