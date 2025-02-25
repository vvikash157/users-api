package db

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var once sync.Once

func InitRedis() {
	once.Do(func() {
		redisHost := os.Getenv("REDIS_HOST")
		redisPort := os.Getenv("REDIS_PORT")
		// redisPassword := os.Getenv("REDIS_PASSWORD")

		redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

		RedisClient = redis.NewClient(&redis.Options{
			Addr:         redisAddr,
			Password:     "",
			DB:           0,
			PoolSize:     10,
			MinIdleConns: 2,
			IdleTimeout:  5 * time.Minute,
		})

		ctx := context.Background()

		_, err := RedisClient.Ping(ctx).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}

		log.Println(" Connected to Redis successfully!")
	})
}

func GetRedisClient() *redis.Client {
	return RedisClient
}

func CloseRedis() {
	if RedisClient != nil {
		err := RedisClient.Close()
		if err != nil {
			log.Printf("Error closing Redis: %v", err)
		}
	}
}
