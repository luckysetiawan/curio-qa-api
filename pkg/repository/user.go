// Package repository stores all database logic the server uses.
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

// userRepo stores mongo client, redis client, and user logic functions.
type userRepo struct {
	mongoClient *mongo.Client
	redisClient *redis.Client
}

// NewUserRepository returns userRepo struct.
func NewUserRepository(mongoClient *mongo.Client, redisClient *redis.Client) IUserRepository {
	return &userRepo{
		mongoClient: mongoClient,
		redisClient: redisClient,
	}
}

// GetLoginStatuses returns all login statuses.
func (r *userRepo) GetLoginStatuses() ([]string, error) {
	keys, err := r.redisClient.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

// CheckUsernameTaken checks whether an username exists or not.
func (r *userRepo) CheckUsernameTaken(username string) bool {
	filter := bson.M{"username": username}
	args := options.FindOne().SetProjection(bson.M{"password": 0})

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

// GetAll returns all user data.
func (r *userRepo) GetAll(filter primitive.M, args ...*options.FindOptions) ([]entity.User, error) {
	coll := r.mongoClient.Database("db").Collection("user")
	var users []entity.User

	opts := options.Find()
	if len(args) > 0 {
		opts = args[0]
	}

	cursor, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var user entity.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Find returns one user with filter.
func (r *userRepo) Find(filter primitive.M, args ...*options.FindOneOptions) (entity.User, error) {
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

// Insert inserts a new user.
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

// MarkLoginStatus marks user login status to redis cache.
func (r *userRepo) MarkLoginStatus(userID string, username string) error {
	expiration := constant.TokenExpirationTime * time.Minute

	err := r.redisClient.Set(context.Background(), userID, username, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// ClearLoginStatus removes user login status from redis cache.
func (r *userRepo) ClearLoginStatus(userID string) error {
	err := r.redisClient.Del(context.Background(), userID).Err()
	if err != nil {
		return err
	}

	return nil
}
