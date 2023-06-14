package handler

import (
	"context"

	core "go-microservices.org/core/proto"
)

// Handler ...
type Handler struct{}

// Login Call is a single request handler called via client.Call or the generated client code
func (*Handler) Login(
	ctx context.Context, in *core.UserRequest, out *core.UserResponse,
) error {
	result, err := GetUsecase().DoLogin(in.User.Username, in.User.Password)
	out.Info = result
	return err
}

// RefreshToken Call is a single request handler called via client.Call or the generated client code
func (*Handler) RefreshToken(
	ctx context.Context, in *core.UserRequest, out *core.UserResponse,
) error {
	result, err := GetUsecase().DoRefreshToken(in.RefreshToken)
	out.Info = result
	return err
}

// Logout Call is a single request handler called via client.Call or the generated client code
func (*Handler) Logout(
	ctx context.Context, in *core.UserRequest, out *core.UserResponse,
) error {
	result, err := GetUsecase().DoLogout(in.AccessToken)
	out.Info = result
	return err
}

// ChangePassword Call is a single request handler called via client.Call or the generated client code
func (*Handler) ChangePassword(
	ctx context.Context, in *core.UserRequest, out *core.UserResponse,
) error {
	result, err := GetUsecase().DoChangePassword(in.User, in.NewPassword)
	out.Info = result
	return err
}

// CreateUser Call is a single request handler called via client.Call or the generated client code
func (*Handler) CreateUser(
	ctx context.Context, in *core.UserRequest, out *core.UserResponse,
) error {
	result, err := GetUsecase().DoCreateUser(in.User)
	out.Info = result
	return err
}

// CreateUserPermission Call is a single request handler called via client.Call or the generated client code
func (*Handler) CreateUserPermission(
	ctx context.Context, in *core.UserRequest, out *core.UserResponse,
) error {
	result, err := GetUsecase().DoCreateUserPermission(in.Permission)
	out.Info = result
	return err
}

// CreateUserRole Call is a single request handler called via client.Call or the generated client code
func (*Handler) CreateUserRole(
	ctx context.Context, in *core.UserRequest, out *core.UserResponse,
) error {
	result, err := GetUsecase().DoCreateUserRole(in.Role)
	out.Info = result
	return err
}
