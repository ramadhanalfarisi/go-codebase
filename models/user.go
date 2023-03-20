package models

import "github.com/google/uuid"

type UserRegister struct {
	Id                   uuid.UUID `json:"id" validate:"required"`
	Email                string    `json:"email" validate:"required"`
	Roles                string    `json:"roles" validate:"required"`
	Password             string    `json:"password" validate:"required"`
	PasswordVerification string    `json:"password_verification" validate:"required" eqfield=Password`
	CreatedAt            string    `json:"created_at"`
	UpdatedAt            string    `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DataLogin struct {
	Token   string `json:"token"`
	Type    string `json:"type"`
	Expired uint32 `json:"expired"`
}
