package usecase

import (
	"fmt"

	auth "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

const (
	tokenParamEmpty = "token param cannot empty"
)

// GetRequestToken ...
func (uc *UseCase) GetRequestToken(tokenParam *auth.TokenData) (*auth.AuthData, error) {
	var (
		aToken, rToken string
		err            error
		errText        = "Auth.GetRequestToken.errCase"
	)

	if tokenParam == nil {
		return nil, utils.SendLogError(errText, fmt.Errorf(tokenParamEmpty))
	}

	if aToken, err = utils.JwtCreateAccessToken(tokenParam); err != nil {
		return nil, utils.SendLogError(errText, err)
	}

	if rToken, err = utils.JwtCreateRefreshToken(tokenParam.UserID); err != nil {
		return nil, utils.SendLogError(errText, err)
	}

	expires := utils.GetAccessTokenExpires()
	result := &auth.AuthData{
		AccessToken:  aToken,
		RefreshToken: rToken,
		ExpireAt:     utils.GetTokenExpiresTime(expires),
		Success:      true,
	}

	return result, nil
}

// GetValidateAccessToken ...
func (uc *UseCase) GetValidateAccessToken(token string) (*auth.AuthData, error) {
	return utils.JwtParseAccessToken(token), nil
}

// GetValidateRefreshToken ...
func (uc *UseCase) GetValidateRefreshToken(refreshToken string) (*auth.AuthData, error) {
	return utils.JwtParseRefreshToken(refreshToken), nil
}
