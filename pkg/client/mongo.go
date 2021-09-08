package client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func NewMongoClient() (*mongo.Database, error) {
	uri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(os.Getenv("MONGO_DB"))
	return db, nil
}
