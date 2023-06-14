package client

import (
	"context"

	coreclient "go-microservices.org/core/client"
	service "go-microservices.org/core/proto"
)

// Client as GRPC client for other service
type Client struct {
	authClient service.AuthService
}

// NewClient ...
func NewClient() *Client {
	return &Client{
		authClient: coreclient.GetAuthService(),
	}
}

// GetRequestAccessTokenSession ...
func (cl *Client) GetRequestAccessTokenSession(tokenParam *service.TokenData) (*service.AuthResponse, error) {
	return cl.authClient.RequestAccessTokenSession(
		context.Background(),
		&service.AuthRequest{
			TokenParam: tokenParam,
		},
	)
}

// GetRequestRefreshTokenSession ...
func (cl *Client) GetRequestRefreshTokenSession(tokenParam *service.TokenData, oldRefreshToken string) (*service.AuthResponse, error) {
	return cl.authClient.RequestRefreshTokenSession(
		context.Background(),
		&service.AuthRequest{
			TokenParam:   tokenParam,
			RefreshToken: oldRefreshToken,
		},
	)
}

// GetValidateRefreshTokenSession ...
func (cl *Client) GetValidateRefreshTokenSession(refreshToken string) (*service.AuthResponse, error) {
	return cl.authClient.ValidateRefreshTokenSession(
		context.Background(),
		&service.AuthRequest{
			RefreshToken: refreshToken,
		},
	)
}

// DoRemoveTokenSession ...
func (cl *Client) DoRemoveTokenSession(accessToken string) (*service.AuthResponse, error) {
	return cl.authClient.RemoveTokenSession(
		context.Background(),
		&service.AuthRequest{
			AccessToken: accessToken,
		},
	)
}
