package repository

import (
	core "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

// GetUserByID ...
func (ar *AbstractRepository) GetUserByID(userID int) (*core.UserData, error) {
	var (
		records [][]string
		err     error
		errText = "User.GetUserByID.err"
	)

	if records, err = ar.GetUserByIDQuery(userID); err != nil {
		return nil, utils.SendLogError(errText+"Query", err)
	}

	result := convertCsvToUser(records)
	return result, nil
}

// GetUserByUsername ...
func (ar *AbstractRepository) GetUserByUsername(username string) (*core.UserData, error) {
	var (
		records [][]string
		err     error
		errText = "User.GetUserByUsername.err"
	)

	if records, err = ar.GetUserByUsernameQuery(username); err != nil {
		return nil, utils.SendLogError(errText+"Query", err)
	}

	result := convertCsvToUser(records)
	return result, nil
}

// GetUserDataByID ...
func (ar *AbstractRepository) GetUserDataByID(userID int) (*core.UserData, error) {
	var (
		records [][]string
		err     error
		errText = "User.GetUserDataByID.err"
	)

	if records, err = ar.GetUserDataByIDQuery(userID); err != nil {
		return nil, utils.SendLogError(errText+"Query", err)
	}

	return convertCsvToUserInfo(records), nil
}

// GetUserDataByUsername ...
func (ar *AbstractRepository) GetUserDataByUsername(username string) (*core.UserData, error) {
	var (
		records [][]string
		err     error
		errText = "User.GetUserDataByUsername.err"
	)

	if records, err = ar.GetUserDataByUsernameQuery(username); err != nil {
		return nil, utils.SendLogError(errText+"Query", err)
	}

	return convertCsvToUserInfo(records), nil
}

// GetUserPermission ...
func (ar *AbstractRepository) GetUserPermission(menu string) (*core.UserPermission, error) {
	var (
		records [][]string
		err     error
		errText = "User.GetUserPermission.err"
	)

	if records, err = ar.GetUserPermissionQuery(menu); err != nil {
		return nil, utils.SendLogError(errText+"Query", err)
	}

	result := convertCsvToUserPermission(records)
	return result, nil
}

// GetUserPermissionBulk ...
func (ar *AbstractRepository) GetUserPermissionBulk(menu []string) ([]*core.UserPermission, error) {
	var (
		records [][]string
		err     error
		errText = "User.GetUserPermissionBulk.err"
	)

	if records, err = ar.GetUserPermissionBulkQuery(menu); err != nil {
		return nil, utils.SendLogError(errText+"Query", err)
	}

	result := convertCsvToUserPermissionBulk(records)
	return result, nil
}

// GetUserRole ...
func (ar *AbstractRepository) GetUserRole(name string) (*core.UserRole, error) {
	var (
		records [][]string
		err     error
		errText = "User.GetUserRole.err"
	)

	if records, err = ar.GetUserRoleQuery(name); err != nil {
		return nil, utils.SendLogError(errText+"Query", err)
	}

	result := convertCsvToUserRole(records)
	return result, nil
}

func convertCsvToUser(records [][]string) *core.UserData {
	if len(records) < 2 {
		return nil
	}

	rec := records[1]
	result := &core.UserData{
		Password: rec[1],
		Username: rec[2],
		Fullname: rec[3],
	}

	result.ID, _ = utils.StringToInt32(rec[0])
	return result
}

func convertCsvToUserInfo(records [][]string) *core.UserData {
	if len(records) < 2 {
		return nil
	}

	rec := records[1]
	result := &core.UserData{
		Password:    rec[1],
		Username:    rec[2],
		Fullname:    rec[3],
		Permissions: buildPermissionsFromJSON(rec[4]),
	}

	result.ID, _ = utils.StringToInt32(rec[0])
	return result
}

func convertCsvToUserPermission(records [][]string) *core.UserPermission {
	if len(records) < 2 {
		return nil
	}

	rec := records[1]
	result := &core.UserPermission{
		Menu:        rec[0],
		Description: rec[1],
	}

	return result
}

func convertCsvToUserPermissionBulk(records [][]string) []*core.UserPermission {
	var result []*core.UserPermission

	for i, rec := range records {
		if i == 0 {
			continue
		}

		perm := &core.UserPermission{
			Menu:        rec[0],
			Description: rec[1],
		}

		result = append(result, perm)
	}

	return result
}

func convertCsvToUserRole(records [][]string) *core.UserRole {
	if len(records) < 2 {
		return nil
	}

	rec := records[1]
	result := &core.UserRole{
		Name:        rec[0],
		Description: rec[1],
		Permissions: buildPermissionsFromJSON(rec[2]),
	}

	return result
}

func buildPermissionsFromJSON(perm string) []*core.UserPermission {
	var permissions []*core.UserPermission
	mapPermission, _ := utils.StringToJSONMap(perm)

	for menu, values := range mapPermission {
		var control []string
		controls, _ := values.([]interface{})

		for _, c := range controls {
			ctr, _ := c.(string)
			control = append(control, ctr)
		}

		permissions = append(
			permissions,
			&core.UserPermission{
				Menu:    menu,
				Control: control,
			},
		)
	}

	return permissions
}
