package usecase

import (
	"strings"

	core "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

const (
	userExist          = "user already exist"
	userNotFound       = "user not found"
	areaNotFound       = "area not found"
	permissionExist    = "permission already exist"
	permissionNotFound = "permission not found"
	roleExist          = "role already exist"
	roleNotFound       = "role not found"
	invalidPassword    = "invalid password"
)

// DoLogin ...
func (uc *UseCase) DoLogin(username, password string) (*core.UserInfo, error) {
	var (
		userData *core.UserData
		userInfo *core.UserInfo
		err      error
		errText  = "User.DoLogin.err"
	)

	if userData, err = uc.repository.GetUserDataByUsername(username); err != nil {
		return nil, err
	} else if userInfo = validateUserPassword(userData, password); !userInfo.Success {
		return userInfo, nil
	}

	if err = uc.doRequestToken(userData, userInfo, ""); err != nil {
		return nil, utils.SendLogError(errText+"Client", err)
	}

	return userInfo, nil
}

// DoRefreshToken ...
func (uc *UseCase) DoRefreshToken(refreshToken string) (*core.UserInfo, error) {
	var (
		userData      *core.UserData
		userInfo      *core.UserInfo
		tokenValidate *core.AuthResponse
		authData      *core.AuthData
		err           error
		errText       = "User.DoRefreshToken.err"
	)

	if tokenValidate, err = uc.client.GetValidateRefreshTokenSession(refreshToken); err != nil {
		return nil, utils.SendLogError(errText+"Client.Validate", err)
	} else if authData = tokenValidate.AuthToken; !authData.Success {
		return getFailedUserInfo(authData.Message), nil
	}

	if userData, err = uc.repository.GetUserDataByID(int(authData.ParsedToken.UserID)); err != nil {
		return nil, err
	} else if userInfo = validateUserExist(userData); !userInfo.Success {
		return userInfo, nil
	}

	if err = uc.doRequestToken(userData, userInfo, refreshToken); err != nil {
		return nil, utils.SendLogError(errText+"Client.Request", err)
	}

	return userInfo, nil
}

// DoLogout ...
func (uc *UseCase) DoLogout(accessToken string) (*core.UserInfo, error) {
	var (
		tokenAuth *core.AuthResponse
		err       error
		errText   = "User.DoLogout.err"
	)

	if tokenAuth, err = uc.client.DoRemoveTokenSession(accessToken); err != nil {
		return nil, utils.SendLogError(errText+"Client", err)
	}

	authData := tokenAuth.AuthToken
	logoutData := &core.UserInfo{
		Message: authData.Message,
		Success: authData.Success,
	}

	return logoutData, nil
}

func (uc *UseCase) doRequestToken(userData *core.UserData, userInfo *core.UserInfo, refreshToken string) error {
	var (
		authData  *core.AuthData
		tokenData = buildTokenParam(userData)
		tokenAuth *core.AuthResponse
		err       error
	)

	if refreshToken != "" {
		tokenAuth, err = uc.client.GetRequestRefreshTokenSession(tokenData, refreshToken)
	} else {
		tokenAuth, err = uc.client.GetRequestAccessTokenSession(tokenData)
	}

	if err != nil {
		return err
	} else if authData = tokenAuth.AuthToken; !authData.Success {
		userInfo.Message = authData.Message
		userInfo.Success = false
		return nil
	}

	userInfo.AccessToken = authData.AccessToken
	userInfo.RefreshToken = authData.RefreshToken
	userInfo.ExpireAt = authData.ExpireAt
	return nil
}

func getFailedUserInfo(message string) *core.UserInfo {
	return &core.UserInfo{
		Message: message,
	}
}

func getSuccessUserInfo(user *core.UserData) *core.UserInfo {
	return &core.UserInfo{
		User:    user,
		Success: true,
	}
}

func validateUserPassword(userData *core.UserData, password string) *core.UserInfo {
	if userData == nil {
		return getFailedUserInfo(userNotFound)
	} else if userData.Password != utils.GetHashPassword(password) {
		return getFailedUserInfo(invalidPassword)
	}

	return getSuccessUserInfo(userData)
}

func validateUserExist(userData *core.UserData) *core.UserInfo {
	if userData == nil {
		return getFailedUserInfo(userNotFound)
	}

	return getSuccessUserInfo(userData)
}

func validatePermissionBulkExist(permission []*core.UserPermission, menu []string) *core.UserInfo {
	if permission == nil {
		return getFailedUserInfo(permissionNotFound)
	}

	menuPerm := buildMenusFromPermissions(permission)
	for _, m := range menu {
		if !utils.FindString(menuPerm, strings.ToLower(m)) {
			return getFailedUserInfo(m + " " + permissionNotFound)
		}
	}

	return getSuccessUserInfo(nil)
}

func buildTokenParam(userData *core.UserData) *core.TokenData {
	tokenData := &core.TokenData{
		UserID: userData.ID,
	}

	utils.CopyObject(userData.Permissions, &tokenData.Permissions)
	return tokenData
}

func buildMenusFromPermissions(permission []*core.UserPermission) []string {
	var menu []string
	for _, perm := range permission {
		menu = append(menu, perm.Menu)
	}

	return menu
}
