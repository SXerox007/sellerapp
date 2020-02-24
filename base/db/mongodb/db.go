package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const (
	DB_NAME        = "sellerapp"
	DB_MONGODB_URL = "mongodb://mongodb:27017"
)

func InitDB() error {
	client, _ = mongo.NewClient(options.Client().ApplyURI(DB_MONGODB_URL))
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err := client.Connect(ctx)
	return err
}

func CreateCollection(collectionName string) *mongo.Collection {
	return client.Database(DB_NAME).Collection(collectionName)
}

func CloseMongoDB() {
	client.Disconnect(context.TODO())
}
