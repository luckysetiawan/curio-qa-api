// Package database provides the database clients to support the functionality
// of the server.
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

// NewMongoClient returns a mongo database client.
func NewMongoClient() *mongo.Client {
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")
	mongoURI := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("MongoDB: ", err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("MongoDB: ", err.Error())
	}

	return client
}
