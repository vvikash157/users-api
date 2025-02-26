package routes

import (
	"Login/controllers"
	"Login/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.UserSignUp).Methods("POST")
	router.Handle("/login", middleware.RateLimitingMiddleware(middleware.AuthMiddleware(http.HandlerFunc(controllers.UserLogin)))).Methods("POST")

	router.Handle("/tasks", middleware.RateLimitingMiddleware(middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateTask)))).Methods("POST")
	router.Handle("/tasks", middleware.RateLimitingMiddleware(middleware.AuthMiddleware(http.HandlerFunc(controllers.GetTasks)))).Methods("GET")
	router.Handle("/tasks/get/{id}", middleware.RateLimitingMiddleware(middleware.AuthMiddleware(http.HandlerFunc(controllers.GetTaskHandler)))).Methods("GET")
	router.Handle("/tasks/{id}", middleware.RateLimitingMiddleware(middleware.AuthMiddleware(http.HandlerFunc(controllers.UpdateTaskHandler)))).Methods("PUT")
	router.Handle("/task/delete/{id}", middleware.RateLimitingMiddleware(middleware.AuthMiddleware(http.HandlerFunc(controllers.DeleteTaskHandler)))).Methods("DELETE")
	router.Handle("/logout", middleware.RateLimitingMiddleware(middleware.AuthMiddleware(http.HandlerFunc(controllers.LogoutHandler)))).Methods("POST")

}
