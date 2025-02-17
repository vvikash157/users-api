package main

import (
	"Login/config"
	"Login/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()
	//config.DB.AutoMigrate(&models.User{})

	router := mux.NewRouter()
	routes.SetupAuthenticationRutes(router)

	fmt.Println("connection tanatan")
	log.Fatal(http.ListenAndServe(":8080", router))
}
