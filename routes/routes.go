package routes

import (
	"Login/controllers"
	"Login/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.UserSignUp).Methods("POST")
	router.Handle("/login", middleware.AuthMiddleware(http.HandlerFunc(controllers.UserLogin))).Methods("POST")

	// protectedRoutes := router.PathPrefix("/dev").Subrouter()
	// protectedRoutes.Use(middleware.AuthMiddleware)

	// protectedRoutes.HandleFunc("/profile", controllers.GetUserProfile).Methods("GET")
	// protectedRoutes.HandleFunc("/update", controllers.UpdateUser).Methods("POST")
}
