package user_controller

import "net/http"

type UserInterface interface {
	RegisterUser(http.ResponseWriter, *http.Request)
	LoginUser(http.ResponseWriter, *http.Request)
}