package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"go-microservices.org/api/graph"
	"go-microservices.org/api/graph/model"
	"go-microservices.org/api/service"
)

func (r *userResolver) Logout(ctx context.Context, obj *model.AbstractModel) (*model.UserResponse, error) {
	authData := initAuthUser(ctx)
	if authData == nil {
		return getSuccessUserResponse(), nil
	}

	return service.DoLogout(authData.AccessToken)
}

func (r *userMutationResolver) ChangePassword(ctx context.Context, obj *model.AbstractModel, currentPassword string, newPassword string, confirmNewPassword string) (*model.UserResponse, error) {
	if success, message := isNewPasswordValid(currentPassword, newPassword, confirmNewPassword); !success {
		return getFailedUserResponse(message), nil
	}

	authData := initAuthUser(ctx)
	if authData == nil {
		return getFailedUserResponse(invalidToken), nil
	}

	return service.DoChangePassword(authData.ParsedToken.UserID, currentPassword, newPassword)
}

func (r *userMutationResolver) Create(ctx context.Context, obj *model.AbstractModel, param model.UserParam) (*model.UserResponse, error) {
	return service.DoCreateUser(param)
}

func (r *userMutationResolver) CreatePermission(ctx context.Context, obj *model.AbstractModel, param model.UserPermissionParam) (*model.UserResponse, error) {
	return service.DoCreateUserPermission(param)
}

func (r *userMutationResolver) CreateRole(ctx context.Context, obj *model.AbstractModel, param model.UserRoleParam) (*model.UserResponse, error) {
	return service.DoCreateUserRole(param)
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

// UserMutation returns graph.UserMutationResolver implementation.
func (r *Resolver) UserMutation() graph.UserMutationResolver { return &userMutationResolver{r} }

type userResolver struct{ *Resolver }
type userMutationResolver struct{ *Resolver }
