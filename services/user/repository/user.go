package repository

import (
	core "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

// DoChangePassword ...
func (ar *AbstractRepository) DoChangePassword(userID int, password string) error {
	var (
		err     error
		errText = "User.DoChangePassword.err"
	)

	if err = ar.UpdatePasswordQuery(userID, password); err != nil {
		return utils.SendLogError(errText+"Query", err)
	}

	return nil
}

// DoCreateUser ...
func (ar *AbstractRepository) DoCreateUser(user *core.UserData) (int, error) {
	var (
		id      int
		err     error
		errText = "User.DoCreateUser.err"
	)

	if id, err = ar.CreateUserQuery(user); err != nil {
		return 0, utils.SendLogError(errText+"Query", err)
	}

	return id, nil
}

// DoCreateUserPermission ...
func (ar *AbstractRepository) DoCreateUserPermission(permission *core.UserPermission) error {
	var (
		err     error
		errText = "User.DoCreateUserPermission.err"
	)

	if err = ar.CreateUserPermissionQuery(permission); err != nil {
		return utils.SendLogError(errText+"Query", err)
	}

	return nil
}

// DoCreateUserRole ...
func (ar *AbstractRepository) DoCreateUserRole(role *core.UserRole) error {
	var (
		err     error
		errText = "User.DoCreateUserRole.err"
	)

	if err = ar.CreateUserRoleQuery(role); err != nil {
		return utils.SendLogError(errText+"Query", err)
	}

	return nil
}
