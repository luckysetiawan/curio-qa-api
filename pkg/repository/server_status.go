package repository

import (
	"context"
	"time"

	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type serverStatusRepo struct {
	mongoClient *mongo.Client
	redisClient *redis.Client
}

func NewServerStatusRepository(mongoClient *mongo.Client, redisClient *redis.Client) IServerStatusRepository {
	return &serverStatusRepo{
		mongoClient: mongoClient,
		redisClient: redisClient,
	}
}

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
