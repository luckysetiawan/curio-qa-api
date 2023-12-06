package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	var (
		mongoHost string
		mongoPort string
		mongoURI  string
		ctx       context.Context
		cancel    context.CancelFunc
		err       error
		client    *mongo.Client
	)

	mongoHost = os.Getenv("MONGO_HOST")
	mongoPort = os.Getenv("MONGO_PORT")
	mongoURI = fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("MongoDB: ", err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("MongoDB: ", err.Error())
	}

	return client
}
