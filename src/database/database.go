package database

import (
	"context"
	"log"
	"time"

	"github.com/WanDmean/graphql-go/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Users *mongo.Collection
	Todos *mongo.Collection
)

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(config.DATABASE)
	/* create collection */
	Users = db.Collection("users")
	Todos = db.Collection("todos")
}
