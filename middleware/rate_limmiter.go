package middleware

import (
	"Login/db"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx                  = context.Background()
	redisRateLimitClient = db.GetRateLimitClient()
)

// RateLimitingMiddleware applies rate limiting based on IP
func RateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		if ip == "" {
			http.Error(w, "Unable to determine IP", http.StatusInternalServerError)
			return
		}

		key := fmt.Sprintf("rate_limit:%s", ip)

		limit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
		if err != nil {
			limit = 10
		}

		window, err := strconv.Atoi(os.Getenv("RATE_WINDOW"))
		if err != nil {
			window = 60
		}

		count, err := redisRateLimitClient.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			log.Println("Redis error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("User IP:", ip, "Current Count:", count)

		if count >= limit {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		// Increment request count and set TTL if it's the first request
		pipe := redisRateLimitClient.TxPipeline()
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, time.Duration(window)*time.Second)
		_, err = pipe.Exec(ctx)
		if err != nil {
			log.Println("Redis pipeline error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getIP(r *http.Request) string {
	// Check for X-Forwarded-For header (comma-separated list)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0]) // First IP is the real client IP
	}

	// Check for X-Real-IP (alternative proxy header)
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fallback to RemoteAddr (includes port), Return as-is if splitting fails
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
