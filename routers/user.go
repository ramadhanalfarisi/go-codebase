package routers

import (
	"github.com/ramadhanalfarisi/go-codebase-kocak/controllers/user"
)

func (r *Router) UserRouter() {
	user_controller := user_controller.NewUserController(r.DB)
	r.Router.HandleFunc("/register", user_controller.RegisterUser).Methods("POST")
	r.Router.HandleFunc("/login", user_controller.LoginUser).Methods("POST")

}
