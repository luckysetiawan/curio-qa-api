package database

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisAddress := fmt.Sprintf("%s:%s", redisHost, redisPort)

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		fmt.Println("Redis: ", err.Error())
	}

	return rdb
}
