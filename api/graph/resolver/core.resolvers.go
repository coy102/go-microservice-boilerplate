package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"go-microservices.org/api/graph"
	"go-microservices.org/api/graph/model"
)

func (r *mutationResolver) User(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

func (r *queryResolver) User(ctx context.Context) (*model.AbstractModel, error) {
	return new(model.AbstractModel), nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
