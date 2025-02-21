package service

import (
	"Login/db"
	m "Login/models"
	"Login/utils"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Logger instance
var log = logrus.New()

// CreateUser inserts a new user into the database
func CreateUser() (map[string]interface{}, error) {
	var u *m.User
	u.UserID = uuid.New().String()

	// Hash the password before storing it
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Error("Error while hashing password: ", err)
		return nil, err
	}
	u.Password = hashedPassword

	// Generate access token
	accessToken, err := utils.GenerateJWT(u.UserID)
	if err != nil {
		log.Error("Error while generating access token: ", err)
		return nil, err
	}
	u.AccessToken = accessToken

	// Insert user into the database
	if err := db.DB.Create(&u).Error; err != nil {
		log.Error("Error while inserting user into DB: ", err)
		return nil, err
	}

	// Return response
	response := map[string]interface{}{
		"access_token": u.AccessToken,
		"user_id":      u.UserID,
	}
	return response, nil
}

// AuthenticateUsers checks user credentials
func AuthenticateUsers(userid, email, password string) (*m.User, error) {
	var foundUser m.User

	// Fetch user by email
	err := db.DB.Where("email = ? AND user_id = ?", email, userid).First(&foundUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		log.Error("Error fetching user: ", err)
		return nil, err
	}

	// Validate password
	if !utils.CheckHashPassword(password, foundUser.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	log.Info("User successfully logged in: ", foundUser.UserID)
	return &foundUser, nil
}
