package controllers

import (
	"Login/models"
	"Login/utils"
	"encoding/json"
	"net/http"
)

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = user.CreateUser()
	if err != nil {
		http.Error(w, "error to create user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created successfully!!"})

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var credentials models.Users
	errr := json.NewDecoder(r.Body).Decode(&credentials)
	if errr != nil {
		http.Error(w, "Invalid user details", http.StatusBadRequest)
		return
	}

	user, err := models.AuthenticateUsers(credentials)
	if err != nil {
		http.Error(w, "Authentication failed for users", http.StatusBadRequest)
		return
	}

	token, erro := utils.GenerateJWT(user.ID)
	if erro != nil {
		http.Error(w, "error while token genereting", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})

}
