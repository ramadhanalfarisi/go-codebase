package repository

import (
	"database/sql"
	"time"

	"github.com/ramadhanalfarisi/go-codebase/helpers"
	"github.com/ramadhanalfarisi/go-codebase/helpers/query_builder"
	user_model "github.com/ramadhanalfarisi/go-codebase/services/user/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}

func (u *UserRepository) InsertUser(model user_model.UserRegisterInput) error {
	query, args := query_builder.New("users").Insert("email", "roles", "password", "created_at").Values(model.Email, model.Roles, model.Password, time.Now()).Build()
	_, err := u.DB.Exec(query, args...)
	if err != nil {
		helpers.Error(err)
		return err
	} else {
		return nil
	}
}

func (u *UserRepository) GetUserByEmail(model user_model.UserLoginInput) (user_model.DataUser, error) {
	var dataLogin user_model.DataUser
	query, args := query_builder.New("users").Select("id", "email", "password", "roles").Where("email = ?", model.Email).Build()
	rows, err := u.DB.Query(query, args...)
	if err != nil {
		helpers.Error(err)
		return user_model.DataUser{}, err
	} else {
		if rows != nil {
			for rows.Next() {
				var (
					id       string
					email    string
					password string
					roles    string
				)
				err := rows.Scan(&id, &email, &password, &roles)
				if err != nil {
					return user_model.DataUser{}, err
				} else {
					dataLogin = user_model.DataUser{Id: id, Email: email, Password: password, Roles: roles}
				}
			}
		} else {
			return user_model.DataUser{}, nil
		}
		return dataLogin, nil
	}
}
