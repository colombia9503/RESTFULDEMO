package routers

import (
	"github.com/colombia9503/RESTful-SqlServer/controllers"
	"github.com/gorilla/mux"
)

//requeres HEADER: Autorization: Beared + TOKEN
func SetAuthRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth/login", controllers.Auth.Login).Methods("POST")
	return router
}
