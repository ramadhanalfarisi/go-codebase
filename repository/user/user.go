package user_repository

import (
	"database/sql"

	"github.com/ramadhanalfarisi/go-codebase-kocak/models"
)

type User struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserInterface {
	return &User{DB: db}
}

func (u User) UserRegister(model models.UserRegister) error {
	return nil
}

func (u User) UserLogin(model models.UserLogin) (models.DataLogin, error) {
	return models.DataLogin{}, nil
}
