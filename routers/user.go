package routers

func (r *Router) UserRouter() {
	r.Router.HandleFunc("/user",r.Controller.RegisterUser).Methods("POST")
}