package utils

import (
	"context"
	"errors"
	"os"
	"time"

	"Login/db"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))
var log = logrus.New()

func GenerateJWT(userID string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userid":     userID,
		"jwt_expiry": time.Now().Add(expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["userid"].(string)
	if !ok {
		return "", errors.New("user_id not found in token")
	}

	return userID, nil
}

func GenerateAndStoreTokens(userID string) (map[string]interface{}, error) {
	accessToken, err := GenerateJWT(userID, 7*24*time.Hour) 
	if err != nil {
		log.Error("Error while generating access token: ", err)
		return nil, err
	}
	key := "accessToken:" + userID
	db.GetRedisClient().Set(context.Background(), key, accessToken, 7*24*time.Hour)

	return map[string]interface{}{
		"access_token": accessToken,
		"userid":       userID,
	}, nil
}
