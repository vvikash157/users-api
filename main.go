package main

import (
	"Login/config"
	"Login/db"
	"Login/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Migrations()
	db.ConnectDB()
	log := config.InitializeLogger()

	//config.DB.AutoMigrate(&models.User{})

	router := mux.NewRouter()
	// s := config.NewGlobalUse(db, log)
	routes.SetupRoutes(router)

	log.Info("connection tanatan")
	log.Fatal(http.ListenAndServe(":8080", router))
}
