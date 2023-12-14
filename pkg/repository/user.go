package repository

import (
	"context"
	"time"

	"github.com/luckysetiawan/curio-qa-api/internal/constant"
	"github.com/luckysetiawan/curio-qa-api/internal/util"
	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	mongoClient *mongo.Client
	redisClient *redis.Client
}

func NewUserRepository(mongoClient *mongo.Client, redisClient *redis.Client) IUserRepository {
	return &userRepo{
		mongoClient: mongoClient,
		redisClient: redisClient,
	}
}

func (r *userRepo) GetLoginStatuses() ([]string, error) {
	keys, err := r.redisClient.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (r *userRepo) CheckUsernameTaken(username string) bool {
	filter := bson.D{{Key: "username", Value: username}}
	args := options.FindOne().SetProjection(bson.D{{Key: "password", Value: 0}})

	_, err := r.Find(filter, args)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// error not caused by "mongo: no documents in result"
			return true
		} else {
			// "mongo: no documents in result"
			return false
		}
	}

	// username taken
	return true
}

func (r *userRepo) Find(filter primitive.D, args ...*options.FindOneOptions) (entity.User, error) {
	coll := r.mongoClient.Database("db").Collection("user")
	var user entity.User

	opts := options.FindOne()
	if len(args) > 0 {
		opts = args[0]
	}

	result := coll.FindOne(context.Background(), filter, opts)
	if result.Err() != nil {
		return entity.User{}, result.Err()
	}

	err := result.Decode(&user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepo) Insert(user entity.User) (interface{}, error) {
	user.Password = util.HashPassword(user.Password)

	coll := r.mongoClient.Database("db").Collection("user")

	res, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	insertedID := res.InsertedID

	return insertedID, nil
}

func (r *userRepo) MarkLoginStatus(userID string, username string) error {
	expiration := constant.TokenExpiryTime * time.Minute

	err := r.redisClient.Set(context.Background(), userID, username, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) ClearLoginStatus(userID string) error {
	err := r.redisClient.Del(context.Background(), userID).Err()
	if err != nil {
		return err
	}

	return nil
}
