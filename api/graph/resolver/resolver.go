package resolver

import (
	"context"

	"go-microservices.org/api/graph/model"
	"go-microservices.org/api/middleware"
	servicemodel "go-microservices.org/core/proto"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

const (
	invalidToken     = "invalid token data, logout first"
	invalidNewPass   = "new password must be different"
	wrongConfirmPass = "confirm new password must be same"
	fieldRequired    = "field is required"
)

func getFailedUserResponse(message string) *model.UserResponse {
	return &model.UserResponse{
		Message: message,
	}
}

func getSuccessUserResponse() *model.UserResponse {
	return &model.UserResponse{
		Success: true,
	}
}

func isNewPasswordValid(currentPassword, newPassword, confirmPassword string) (bool, string) {
	if newPassword == currentPassword {
		return false, invalidNewPass
	}

	if newPassword != confirmPassword {
		return false, wrongConfirmPass
	}

	return true, ""
}

func initAuthUser(ctx context.Context) *servicemodel.AuthData {
	authData := middleware.GetTokenContext(ctx)
	if authData == nil {
		return nil
	}

	if authData.ParsedToken == nil {
		return nil
	}

	return authData
}
