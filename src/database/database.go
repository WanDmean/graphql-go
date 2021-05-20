package database

import (
	"context"
	"log"
	"time"

	"github.com/WanDmean/graphql-go/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}

func GetCollection(name string) *mongo.Collection {
	collection := Client.Database(config.DATABASE).Collection(name)
	return collection
}
