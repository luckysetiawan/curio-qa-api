package repository

import (
	"context"

	"github.com/luckysetiawan/curio-qa-api/pkg/entity"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type curioRepo struct {
	mongoClient *mongo.Client
	redisClient *redis.Client
}

func NewCurioRepository(mongoClient *mongo.Client, redisClient *redis.Client) ICurioRepository {
	return &curioRepo{
		mongoClient: mongoClient,
		redisClient: redisClient,
	}
}

func (r *curioRepo) Find(userID, curioID primitive.ObjectID) (entity.Curio, error) {
	var (
		user  entity.User
		curio entity.Curio
	)

	coll := r.mongoClient.Database("db").Collection("user")
	filter := bson.M{"_id": userID, "curios._id": curioID}

	err := coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return entity.Curio{}, err
	}

	for _, c := range user.Curios {
		if c.ID == curioID {
			curio = c
			break
		}
	}

	return curio, nil
}

func (r *curioRepo) Insert(userID primitive.ObjectID, curio entity.Curio) error {
	coll := r.mongoClient.Database("db").Collection("user")
	filter := bson.M{"_id": userID}
	update := bson.M{"$push": bson.M{"curios": curio}}

	_, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *curioRepo) UpdateStatus(userID, curioID primitive.ObjectID, status bool) error {
	coll := r.mongoClient.Database("db").Collection("user")
	filter := bson.M{"_id": userID, "curios._id": curioID}
	update := bson.M{"$set": bson.M{"curios.$.status": status}}

	_, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
