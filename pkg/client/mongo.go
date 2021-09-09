package client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func NewMongoClient() (*mongo.Database, error) {
	uri := os.Getenv("MONGO_URI")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancelFunc()
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	db := client.Database(os.Getenv("MONGO_DB"))
	return db, nil
}
