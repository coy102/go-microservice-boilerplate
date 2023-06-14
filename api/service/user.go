package service

import (
	"context"

	"go-microservices.org/api/graph/model"
	coreclient "go-microservices.org/core/client"
	servicemodel "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

// DoLogout ...
func DoLogout(accessToken string) (result *model.UserResponse, err error) {
	service := coreclient.GetUserService()
	response, err := service.Logout(
		context.Background(),
		&servicemodel.UserRequest{
			AccessToken: accessToken,
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.UserResponse)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}

// DoChangePassword ...
func DoChangePassword(userID int32, currentPassword, newPassword string) (result *model.UserResponse, err error) {
	service := coreclient.GetUserService()
	response, err := service.ChangePassword(
		context.Background(),
		&servicemodel.UserRequest{
			User: &servicemodel.UserData{
				ID:       userID,
				Password: currentPassword,
			},
			NewPassword: newPassword,
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.UserResponse)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}

// DoCreateUser ...
func DoCreateUser(param model.UserParam) (result *model.UserResponse, err error) {
	userData := new(servicemodel.UserData)
	if err := utils.CopyObject(param, &userData); err != nil {
		return nil, err
	}

	service := coreclient.GetUserService()
	response, err := service.CreateUser(
		context.Background(),
		&servicemodel.UserRequest{
			User: userData,
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.UserResponse)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}

// DoCreateUserPermission ...
func DoCreateUserPermission(param model.UserPermissionParam) (result *model.UserResponse, err error) {
	userPerm := new(servicemodel.UserPermission)
	if err := utils.CopyObject(param, &userPerm); err != nil {
		return nil, err
	}

	service := coreclient.GetUserService()
	response, err := service.CreateUserPermission(
		context.Background(),
		&servicemodel.UserRequest{
			Permission: userPerm,
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.UserResponse)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}

// DoCreateUserRole ...
func DoCreateUserRole(param model.UserRoleParam) (result *model.UserResponse, err error) {
	userRole := new(servicemodel.UserRole)
	if err := utils.CopyObject(param, &userRole); err != nil {
		return nil, err
	}

	service := coreclient.GetUserService()
	response, err := service.CreateUserRole(
		context.Background(),
		&servicemodel.UserRequest{
			Role: userRole,
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.UserResponse)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}
