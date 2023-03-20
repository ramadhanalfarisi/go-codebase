package user_service

import (
	"database/sql"

	"github.com/ramadhanalfarisi/go-codebase-kocak/models"
)

type User struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) UserInterface {
	return &User{DB: db}
}

func (u User) UserRegister(model models.UserRegister) {

}
func (u User) UserLogin(model models.UserLogin) {

}
