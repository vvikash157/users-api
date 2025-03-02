package middleware

import (
	"Login/config"
	"Login/db"
	"Login/utils"
	"context"
	"net/http"
	"strings"
)

var (
	log         = config.InitializeLogger()
	redisClient = db.GetSessionClient()
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
		ctx := context.Background()

		userID, err := utils.ValidateJWT(token)
		if err != nil {
			log.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		key := "accessToken:" + userID

		sessionExists, err := redisClient.Exists(ctx, key).Result()
		if err != nil || sessionExists == 0 {
			http.Error(w, "Session expired, please log in again", http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		log.Info("user validated with userid : ", userID)

		ctx = context.WithValue(r.Context(), "userid", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
