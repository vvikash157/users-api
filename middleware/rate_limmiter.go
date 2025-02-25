package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func RateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // Use IP address as identifier
		key := fmt.Sprintf("rate_limit:%s", ip)

		count, err := redisClient.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			log.Println("Redis error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		requestCount, _ := strconv.Atoi(count)

		rateLimitPerMinutes := 10

		if requestCount >= rateLimitPerMinutes {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		redisClient.Incr(ctx, key)
		redisClient.Expire(ctx, key, time.Minute)

		next.ServeHTTP(w, r)
	})
}
