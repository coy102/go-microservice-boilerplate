package resolver

import (
	"context"

	"go-microservices.org/api/login/model"
	coreclient "go-microservices.org/core/client"
	servicemodel "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// DoLogin ...
func DoLogin(username, password string) (result *model.LoginData, err error) {
	service := coreclient.GetUserService()
	response, err := service.Login(
		context.Background(),
		&servicemodel.UserRequest{
			User: &servicemodel.UserData{
				Username: username,
				Password: password,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.LoginData)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}

// DoRefreshToken ...
func DoRefreshToken(refreshToken string) (result *model.LoginData, err error) {
	service := coreclient.GetUserService()
	response, err := service.RefreshToken(
		context.Background(),
		&servicemodel.UserRequest{
			RefreshToken: refreshToken,
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.LoginData)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}
