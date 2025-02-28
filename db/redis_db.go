package db

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	SessionClient      *redis.Client
	RateLimitClient    *redis.Client
	NotificationClient *redis.Client
	once              sync.Once
)

func InitRedis() {
	once.Do(func() {
		redisHost := os.Getenv("REDIS_HOST")
		redisPort := os.Getenv("REDIS_PORT")
		redisPassword := os.Getenv("REDIS_PASSWORD")

		redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

		SessionClient = redis.NewClient(&redis.Options{
			Addr:         redisAddr,
			Password:     redisPassword,
			DB:           0,
			PoolSize:     10,
			MinIdleConns: 2,
			IdleTimeout:  5 * time.Minute,
		})

		RateLimitClient = redis.NewClient(&redis.Options{
			Addr:         redisAddr,
			Password:     redisPassword,
			DB:           1,
			PoolSize:     10,
			MinIdleConns: 2,
			IdleTimeout:  5 * time.Minute,
		})

		NotificationClient = redis.NewClient(&redis.Options{
			Addr:         redisAddr,
			Password:     redisPassword,
			DB:           2,
			PoolSize:     10,
			MinIdleConns: 2,
			IdleTimeout:  5 * time.Minute,
		})

		ctx := context.Background()
		if _, err := SessionClient.Ping(ctx).Result(); err != nil {
			log.Fatalf("Failed to connect to Redis Session DB: %v", err)
		}

		if _, err := RateLimitClient.Ping(ctx).Result(); err != nil {
			log.Fatalf("Failed to connect to Redis Rate Limiting DB: %v", err)
		}

		if _, err := NotificationClient.Ping(ctx).Result(); err != nil {
			log.Fatalf("Failed to connect to Redis Notifications DB: %v", err)
		}

		log.Println(" Connected to Redis successfully!")
	})
}

func GetSessionClient() *redis.Client {
	return SessionClient
}

func GetRateLimitClient() *redis.Client {
	return RateLimitClient
}

func GetNotificationClient() *redis.Client {
	return NotificationClient
}

func CloseRedis() {
	if SessionClient != nil {
		err:=SessionClient.Close()
		if err != nil {
			log.Printf("Error closing Session Client Redis: %v", err)
		}
	}
	if RateLimitClient != nil {
		err:=RateLimitClient.Close()
		if err != nil {
			log.Printf("Error closing Rate Limit Client Redis: %v", err)
		}
	}
	if NotificationClient != nil {
		err:=NotificationClient.Close()
		if err != nil {
			log.Printf("Error closing Notification Client Redis: %v", err)
		}
	}
}