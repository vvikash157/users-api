package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("delhiMumbaiChennaiKolkata")

func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userID,
		"jwt_expiry": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
