package service

import (
	"context"

	coreclient "go-microservices.org/core/client"
	servicemodel "go-microservices.org/core/proto"
)

// GetValidateTokenSession ...
func GetValidateTokenSession(token string) (*servicemodel.AuthData, error) {
	service := coreclient.GetAuthService()
	response, err := service.ValidateAccessTokenSession(
		context.Background(),
		&servicemodel.AuthRequest{
			AccessToken: token,
		},
	)

	if err != nil {
		return nil, err
	}

	return response.AuthToken, nil
}
