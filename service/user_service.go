package service

import (
	"Login/db"
	"Login/models"
	"Login/utils"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Logger instance
var log = logrus.New()

func CreateUser(u models.User) (map[string]interface{}, error) {
	u.UserID = uuid.New().String()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Error("Error while hashing password: ", err)
		return nil, err
	}
	u.Password = hashedPassword

	tokens, err := utils.GenerateAndStoreTokens(u.UserID)
	if err != nil {
		return nil, err
	}
	u.AccessToken = tokens["access_token"].(string)

	if err := db.DB.Create(&u).Error; err != nil {
		log.Error("Error while inserting user into DB: ", err)
		return nil, err
	}

	return tokens, nil
}

func AuthenticateUsers(email, password string) (map[string]interface{}, error) {
	var foundUser models.User
	err := db.DB.Where("email = ?", email).First(&foundUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		log.Error("Error fetching user: ", err)
		return nil, err
	}

	if !utils.CheckHashPassword(password, foundUser.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	tokens, err := utils.GenerateAndStoreTokens(foundUser.UserID)
	if err != nil {
		return nil, err
	}

	log.Info("User successfully logged in: ", foundUser.UserID)
	return tokens, nil
}
