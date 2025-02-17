package routes

import (
	"Login/controllers"

	"github.com/gorilla/mux"
)

func SetupAuthenticationRutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.UserSignUp).Methods("POST")
	router.HandleFunc("/login", controllers.UserLogin).Methods("POST")
}
