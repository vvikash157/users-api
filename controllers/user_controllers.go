package controllers

import (
	"Login/db"
	"Login/models"
	"Login/service"
	"encoding/json"
	"net/http"
)

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	userDetails, err := service.CreateUser()
	if err != nil {
		http.Error(w, "error to create user", http.StatusBadRequest)
		return
	}

	userDetails["message"] = "user created successfully"
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userDetails)

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value("userid").(string)

	if userID == "" {
		http.Error(w, "Unauthorized: No user ID found", http.StatusUnauthorized)
		return
	}
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid user details", http.StatusBadRequest)
		return
	}

	user, err := service.AuthenticateUsers(userID, credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, "Authentication failed for users", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"user_id": user.UserID,
		"message": " user successfully logged in the system!! ",
	}
	json.NewEncoder(w).Encode(response)

}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from context
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var user models.User
	err := db.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return user profile (excluding password)
	response := map[string]interface{}{
		"id":      user.ID,
		"user_id": user.UserID,
		"name":    user.Name,
		"email":   user.Email,
		"age":     user.Age,
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find and update the user
	var user models.User
	err := db.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Only update fields that are provided in the request
	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Age != 0 {
		user.Age = updatedUser.Age
	}

	db.DB.Save(&user)

	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}
