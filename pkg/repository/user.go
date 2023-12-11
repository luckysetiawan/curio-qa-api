package repository

import (
	"context"

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

func (r *userRepo) CheckUsernameTaken(username string) bool {
	_, err := r.Find(bson.D{{Key: "username", Value: username}})
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

func (r *userRepo) Find(filter primitive.D) (entity.User, error) {
	coll := r.mongoClient.Database("db").Collection("user")
	opts := options.FindOne().SetProjection(bson.D{{Key: "password", Value: 0}})
	var user entity.User

	res := coll.FindOne(context.TODO(), filter, opts)

	err := res.Decode(user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepo) Insert(user entity.User) (interface{}, error) {
	user.Password = util.HashPassword(user.Password)

	coll := r.mongoClient.Database("db").Collection("user")

	res, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	insertedID := res.InsertedID

	return insertedID, nil
}
