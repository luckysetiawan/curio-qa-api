// Package database provides the database clients to support the functionality
// of the server.
package database

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient returns a redis database client.
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
