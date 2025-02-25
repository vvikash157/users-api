package middleware

import (
	"Login/db"
	"Login/utils"
	"context"
	"fmt"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]
		redisClient := db.GetRedisClient() 
		ctx := context.Background()
		fmt.Println("redis: ", db.GetRedisClient())
		
		userID, err := utils.ValidateJWT(token)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		key := "accessToken:" + userID
	
		sessionExists, err := redisClient.Exists(ctx, key).Result()
		if err != nil || sessionExists == 0 {
			http.Error(w, "Session expired, please log in again", http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		
		ctx = context.WithValue(r.Context(), "userid", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
