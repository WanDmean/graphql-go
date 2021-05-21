package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/WanDmean/graphql-go/graph/generated"
	"github.com/WanDmean/graphql-go/graph/model"
	"github.com/WanDmean/graphql-go/src/pkg/users"
)

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return users.Save(ctx, input)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return users.FindById(ctx, id)
}

func (r *queryResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
