package handler

import (
	"context"

	auth "go-microservices.org/core/proto"
)

// Handler ...
type Handler struct{}

// RequestToken Call is a single request handler called via client.Call or the generated client code
func (*Handler) RequestToken(
	ctx context.Context, in *auth.AuthRequest, out *auth.AuthResponse,
) error {
	result, err := GetUsecase().GetRequestToken(in.TokenParam)
	out.AuthToken = result
	return err
}

// ValidateAccessToken Call is a single request handler called via client.Call or the generated client code
func (*Handler) ValidateAccessToken(
	ctx context.Context, in *auth.AuthRequest, out *auth.AuthResponse,
) error {
	result, err := GetUsecase().GetValidateAccessToken(in.AccessToken)
	out.AuthToken = result
	return err
}

// ValidateRefreshToken Call is a single request handler called via client.Call or the generated client code
func (*Handler) ValidateRefreshToken(
	ctx context.Context, in *auth.AuthRequest, out *auth.AuthResponse,
) error {
	result, err := GetUsecase().GetValidateRefreshToken(in.RefreshToken)
	out.AuthToken = result
	return err
}

// RequestAccessTokenSession Call is a single request handler called via client.Call or the generated client code
func (*Handler) RequestAccessTokenSession(
	ctx context.Context, in *auth.AuthRequest, out *auth.AuthResponse,
) error {
	result, err := GetUsecase().GetRequestAccessTokenSession(in.TokenParam)
	out.AuthToken = result
	return err
}

// RequestRefreshTokenSession Call is a single request handler called via client.Call or the generated client code
func (*Handler) RequestRefreshTokenSession(
	ctx context.Context, in *auth.AuthRequest, out *auth.AuthResponse,
) error {
	result, err := GetUsecase().GetRequestRefreshTokenSession(in.TokenParam, in.RefreshToken)
	out.AuthToken = result
	return err
}

// ValidateAccessTokenSession Call is a single request handler called via client.Call or the generated client code
func (*Handler) ValidateAccessTokenSession(
	ctx context.Context, in *auth.AuthRequest, out *auth.AuthResponse,
) error {
	result, err := GetUsecase().GetValidateAccessTokenSession(in.AccessToken)
	out.AuthToken = result
	return err
}

// ValidateRefreshTokenSession Call is a single request handler called via client.Call or the generated client code
func (*Handler) ValidateRefreshTokenSession(
	ctx context.Context, in *auth.AuthRequest, out *auth.AuthResponse,
) error {
	result, err := GetUsecase().GetValidateRefreshTokenSession(in.RefreshToken)
	out.AuthToken = result
	return err
}

// RemoveTokenSession Call is a single request handler called via client.Call or the generated client code
func (*Handler) RemoveTokenSession(
	ctx context.Context, in *auth.AuthRequest, out *auth.AuthResponse,
) error {
	result, err := GetUsecase().DoRemoveTokenSession(in.AccessToken)
	out.AuthToken = result
	return err
}
