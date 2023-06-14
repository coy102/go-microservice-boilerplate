package postgres

import (
	"encoding/json"

	core "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"
)

// UpdatePasswordQuery ...
func (pr *Repository) UpdatePasswordQuery(userID int, password string) error {
	qapi := "update users set password = $1 where id = $2"
	_, err := pr.Connection.ExecuteDB(qapi, password, userID)
	return err
}

// CreateUserQuery ...
func (pr *Repository) CreateUserQuery(user *core.UserData) (int, error) {
	qapi := `INSERT INTO 
	users (
		password, username, fullname, role_name
	) VALUES (
		$1, $2, $3, $4
	) RETURNING id`

	result, err := pr.Connection.ExecuteDBInsertReturnID(
		qapi,
		utils.GetHashPassword(user.Password),
		user.Username,
		user.Fullname,
		utils.TrimLower(user.RoleName),
	)

	return int(result.InsertedID), err
}

// CreateUserPermissionQuery ...
func (pr *Repository) CreateUserPermissionQuery(permission *core.UserPermission) error {
	qapi := `INSERT INTO 
	user_permission (
		menu, description
	) VALUES (
		$1, $2
	)`

	_, err := pr.Connection.ExecuteDB(
		qapi,
		utils.TrimLower(permission.Menu),
		permission.Description,
	)

	return err
}

// CreateUserRoleQuery ...
func (pr *Repository) CreateUserRoleQuery(role *core.UserRole) error {
	qapi := `INSERT INTO 
	user_role (
		name, description, permissions
	) VALUES (
		$1, $2, $3
	)`

	mapPermission := make(map[string][]string)
	for _, perm := range role.Permissions {
		mapPermission[perm.Menu] = perm.Control
	}

	jsonPermission, _ := json.Marshal(mapPermission)

	_, err := pr.Connection.ExecuteDB(
		qapi,
		utils.TrimLower(role.Name),
		role.Description,
		utils.TrimLower(string(jsonPermission)),
	)

	return err
}
