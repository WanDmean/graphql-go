package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/WanDmean/graphql-go/graph/model"
	"github.com/WanDmean/graphql-go/src/pkg/todos"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return todos.Save(ctx, input), nil
}

func (r *queryResolver) Todos(ctx context.Context, userID string) ([]*model.Todo, error) {
	return todos.All(ctx, userID), nil
}
