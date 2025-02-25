package main

import (
	"Login/config"
	"Login/db"
	"Login/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Migrations()
	db.ConnectDB()
	db.InitRedis()
	defer db.CloseRedis()
	log := config.InitializeLogger()
	router := mux.NewRouter()
	// s := config.NewGlobalUse(db, log)
	routes.SetupRoutes(router)

	log.Info("connection successful")
	log.Fatal(http.ListenAndServe(":8080", router))
}
