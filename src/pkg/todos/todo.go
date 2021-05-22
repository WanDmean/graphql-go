package todos

import (
	"context"
	"log"

	"github.com/WanDmean/graphql-go/graph/model"
	"github.com/WanDmean/graphql-go/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Save(ctx context.Context, input model.NewTodo) *model.Todo {
	collection := database.GetCollection("todos")
	UserObjectID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		log.Fatal(err)
	}
	res, err := collection.InsertOne(ctx, &TodoType{
		UserID: UserObjectID,
		Done:   input.Done,
		Text:   input.Text,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &model.Todo{
		ID:     res.InsertedID.(primitive.ObjectID).Hex(),
		Text:   input.Text,
		Done:   input.Done,
		UserID: input.UserID,
	}
}

func All(ctx context.Context, userID string) []*model.Todo {
	collection := database.GetCollection("todos")
	UserObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Fatal(err)
	}
	cur, err := collection.Find(ctx, &TodoType{
		UserID: UserObjectID,
	})
	if err != nil {
		log.Fatal(err)
	}
	var todos []*model.Todo
	for cur.Next(ctx) {
		todo := TodoType{}
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, &model.Todo{
			ID:     todo.ID.Hex(),
			Text:   todo.Text,
			Done:   todo.Done,
			UserID: todo.UserID.Hex(),
		})
	}
	return todos
}
