package usecase

import (
	core "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

// DoChangePassword ...
func (uc *UseCase) DoChangePassword(user *core.UserData, newPassword string) (*core.UserInfo, error) {
	var (
		userData *core.UserData
		userInfo *core.UserInfo
		err      error
	)

	if userData, err = uc.repository.GetUserByID(int(user.ID)); err != nil {
		return nil, err
	} else if userInfo = validateUserPassword(userData, user.Password); !userInfo.Success {
		return userInfo, nil
	}

	if err = uc.repository.DoChangePassword(int(userData.ID), utils.GetHashPassword(newPassword)); err != nil {
		return nil, err
	}

	return userInfo, nil
}

// DoCreateUser ...
func (uc *UseCase) DoCreateUser(user *core.UserData) (*core.UserInfo, error) {
	if userData, err := uc.repository.GetUserByUsername(user.Username); err != nil {
		return nil, err
	} else if userData != nil {
		return getFailedUserInfo(userExist), nil
	}

	if userRole, err := uc.repository.GetUserRole(user.RoleName); err != nil {
		return nil, err
	} else if userRole == nil {
		return getFailedUserInfo(roleNotFound), nil
	}

	if _, err := uc.repository.DoCreateUser(user); err != nil {
		return nil, err
	}

	return getSuccessUserInfo(nil), nil
}

// DoCreateUserPermission ...
func (uc *UseCase) DoCreateUserPermission(permission *core.UserPermission) (*core.UserInfo, error) {
	if userPerm, err := uc.repository.GetUserPermission(permission.Menu); err != nil {
		return nil, err
	} else if userPerm != nil {
		return getFailedUserInfo(permissionExist), nil
	}

	if err := uc.repository.DoCreateUserPermission(permission); err != nil {
		return nil, err
	}

	return getSuccessUserInfo(nil), nil
}

// DoCreateUserRole ...
func (uc *UseCase) DoCreateUserRole(role *core.UserRole) (*core.UserInfo, error) {
	if userRole, err := uc.repository.GetUserRole(role.Name); err != nil {
		return nil, err
	} else if userRole != nil {
		return getFailedUserInfo(roleExist), nil
	}

	menu := buildMenusFromPermissions(role.Permissions)
	if userPerm, err := uc.repository.GetUserPermissionBulk(menu); err != nil {
		return nil, err
	} else if userInfo := validatePermissionBulkExist(userPerm, menu); !userInfo.Success {
		return userInfo, nil
	}

	if err := uc.repository.DoCreateUserRole(role); err != nil {
		return nil, err
	}

	return getSuccessUserInfo(nil), nil
}
