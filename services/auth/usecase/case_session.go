package usecase

import (
	auth "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

const (
	sessionExpire = "session is not found or already expired"
	active        = "active"
	inactive      = "inactive"
)

// GetRequestAccessTokenSession ...
func (uc *UseCase) GetRequestAccessTokenSession(tokenParam *auth.TokenData) (*auth.AuthData, error) {
	var (
		authData *auth.AuthData
		err      error
	)

	if authData, err = uc.buildTokenSession(tokenParam); err != nil {
		return nil, utils.SendLogError("AccessToken", err)
	}

	return authData, nil
}

// GetRequestRefreshTokenSession ...
func (uc *UseCase) GetRequestRefreshTokenSession(tokenParam *auth.TokenData, oldRefreshToken string) (*auth.AuthData, error) {
	var (
		authData *auth.AuthData
		err      error
	)

	if authData, err = uc.buildTokenSession(tokenParam); err != nil {
		return nil, utils.SendLogError("RefreshToken", err)
	}

	sessionKey := uc.session.CreateKey(oldRefreshToken)
	uc.session.DeleteKey(sessionKey)

	return authData, nil
}

func (uc *UseCase) buildTokenSession(tokenParam *auth.TokenData) (*auth.AuthData, error) {
	var (
		authData *auth.AuthData
		err      error
	)

	if authData, err = uc.GetRequestToken(tokenParam); err != nil {
		return nil, err
	}

	sessionKey := uc.session.CreateKey(authData.AccessToken)
	if err = uc.session.SetKey(sessionKey, active); err != nil {
		return nil, err
	}

	sessionKeyRefresh := uc.session.CreateKey(authData.RefreshToken)
	if err = uc.session.SetKey(sessionKeyRefresh, active); err != nil {
		return nil, err
	}

	authData.SessionKey = sessionKey
	return authData, nil
}

// GetValidateAccessTokenSession ...
func (uc *UseCase) GetValidateAccessTokenSession(accessToken string) (*auth.AuthData, error) {
	var (
		authData *auth.AuthData
		err      error
		errText  = "Auth.GetValidateTokenSession.errCase"
	)

	if authData, err = uc.GetValidateAccessToken(accessToken); err != nil {
		return nil, err
	} else if !authData.Success {
		return authData, nil
	}

	if err = uc.validateResponse(accessToken, authData); err != nil {
		return nil, utils.SendLogError(errText, err)
	}

	return authData, nil
}

// GetValidateRefreshTokenSession ...
func (uc *UseCase) GetValidateRefreshTokenSession(refreshToken string) (*auth.AuthData, error) {
	var (
		authData *auth.AuthData
		err      error
		errText  = "Auth.GetValidateRefreshTokenSession.errCase"
	)

	if authData, err = uc.GetValidateRefreshToken(refreshToken); err != nil {
		return nil, err
	} else if !authData.Success {
		return authData, nil
	}

	if err = uc.validateResponse(refreshToken, authData); err != nil {
		return nil, utils.SendLogError(errText, err)
	}

	return authData, nil
}

func (uc *UseCase) validateResponse(token string, authData *auth.AuthData) error {
	authData.Success = false
	sessionKey := uc.session.CreateKey(token)
	if _, exist, err := uc.session.GetKeyExist(sessionKey); err != nil {
		return err
	} else if !exist {
		authData.Message = sessionExpire
	} else {
		authData.Success = true
	}

	return nil
}

// DoRemoveTokenSession ...
func (uc *UseCase) DoRemoveTokenSession(accessToken string) (*auth.AuthData, error) {
	var (
		authData *auth.AuthData
		err      error
		errText  = "Auth.DoRemoveTokenSession.errCase"
	)

	sessionKey := uc.session.CreateKey(accessToken)
	if err = uc.session.DeleteKey(sessionKey); err != nil {
		return nil, utils.SendLogError(errText+".DeleteKey", err)
	}

	authData = &auth.AuthData{
		SessionKey: sessionKey,
		Success:    true,
	}

	return authData, nil
}
