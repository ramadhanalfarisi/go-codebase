package user_controller

import (
	"database/sql"
	"net/http"
)

type User struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) UserInterface {
	return &User{DB: db}
}

func (u User) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// type some code
}

func (u User) LoginUser(w http.ResponseWriter, r *http.Request) {
	// type some code
}
