package middleware

import (
	"Login/utils"
	"context"
	"fmt"
	"net/http"
	"strings"
)

// AuthMiddleware validates JWT before allowing access
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println("yes middleware executed")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		userID, err := utils.ValidateJWT(tokenParts[1])
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Store userID in request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userid", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
