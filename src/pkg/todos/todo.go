package todos

import (
	"context"
	"fmt"

	"github.com/WanDmean/graphql-go/graph/model"
	"github.com/WanDmean/graphql-go/src/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Save(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	UserObjectID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		return &model.Todo{}, fmt.Errorf("invalid object id")
	}
	res, err := database.Todos.InsertOne(ctx, &TodoType{
		UserID: UserObjectID,
		Done:   input.Done,
		Text:   input.Text,
	})
	if err != nil {
		return &model.Todo{}, err
	}
	return &model.Todo{
		ID:     res.InsertedID.(primitive.ObjectID).Hex(),
		Text:   input.Text,
		Done:   input.Done,
		UserID: input.UserID,
	}, nil
}

func All(ctx context.Context, userID string) ([]*model.Todo, error) {
	UserObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return []*model.Todo{}, fmt.Errorf("invalid object id")
	}
	cur, err := database.Todos.Find(ctx, &TodoType{
		UserID: UserObjectID,
	})
	if err != nil {
		return []*model.Todo{}, err
	}
	var todos []*model.Todo
	for cur.Next(ctx) {
		todo := TodoType{}
		err := cur.Decode(&todo)
		if err != nil {
			return todos, err
		}
		todos = append(todos, &model.Todo{
			ID:     todo.ID.Hex(),
			Text:   todo.Text,
			Done:   todo.Done,
			UserID: todo.UserID.Hex(),
		})
	}
	return todos, nil
}
