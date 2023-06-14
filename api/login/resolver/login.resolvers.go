package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"go-microservices.org/api/login"
	"go-microservices.org/api/login/model"
)

func (r *mutationResolver) Login(ctx context.Context, userID string, password string, refreshToken string) (*model.LoginData, error) {
	if refreshToken != "" {
		return DoRefreshToken(refreshToken)
	}

	return DoLogin(userID, password)
}

func (r *queryResolver) Noop(ctx context.Context) (bool, error) {
	return true, nil
}

// Mutation returns login.MutationResolver implementation.
func (r *Resolver) Mutation() login.MutationResolver { return &mutationResolver{r} }

// Query returns login.QueryResolver implementation.
func (r *Resolver) Query() login.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
