package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/WanDmean/graphql-go/graph/model"
	"github.com/WanDmean/graphql-go/src/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	collection := database.GetCollection("todos")
	UserObjectID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		log.Fatal(err)
	}
	res, err := collection.InsertOne(ctx, bson.D{
		{Key: "text", Value: input.Text},
		{Key: "done", Value: input.Done},
		{Key: "_user", Value: UserObjectID},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &model.Todo{
		ID:     res.InsertedID.(primitive.ObjectID).Hex(),
		Text:   input.Text,
		Done:   input.Done,
		UserID: input.UserID,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*model.Todo, error) {
	collection := database.GetCollection("todos")
	UserObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Fatal(err)
	}
	cur, err := collection.Find(ctx, bson.D{
		{Key: "_user", Value: UserObjectID},
	})
	if err != nil {
		log.Fatal(err)
	}
	var todos []*model.Todo
	for cur.Next(ctx) {
		var todo *model.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
