package models

type UserRegisterInput struct {
	Email    string `json:"email" validate:"required"`
	Roles    string `json:"roles" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DataUser struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
}
