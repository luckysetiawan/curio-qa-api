// Package repository stores all database logic the server uses.
package repository

import (
	"context"
	"time"

	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// serverStatusRepo stores mongo client, redis client, and server status logic
// functions.
type serverStatusRepo struct {
	mongoClient *mongo.Client
	redisClient *redis.Client
}

// NewServerStatusRepository returns serverStatusRepo struct.
func NewServerStatusRepository(mongoClient *mongo.Client, redisClient *redis.Client) IServerStatusRepository {
	return &serverStatusRepo{
		mongoClient: mongoClient,
		redisClient: redisClient,
	}
}

// GetServerStatus returns server status.
func (r *serverStatusRepo) GetServerStatus() entity.ServerStatus {
	var serverStatus entity.ServerStatus

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.mongoClient.Ping(ctx, nil)
	if err == nil {
		serverStatus.MongoStatus = true
	}

	err = r.redisClient.Ping(ctx).Err()
	if err == nil {
		serverStatus.RedisStatus = true
	}

	return serverStatus
}
