package postgres

import (
	"fmt"
	"strconv"
	"strings"

	"go-microservices.org/core/connection"
	"go-microservices.org/core/utils"
)

// Repository ...
type Repository struct {
	Connection connection.Connection
}

const (
	typID int = iota
	typUser
)

func userInfoQuery(paramTyp int) string {
	var where string
	switch paramTyp {
	case typID:
		where = "id"
	case typUser:
		where = "lower(u.username)"
	}

	return fmt.Sprintf(
		`select u.id, u.password, u.username, u.fullname, lower(permissions::varchar) as permission 
		from users u
		inner join user_role ur on ur.name = u.role_name 
		where u.is_active = true and %s = $1
		limit 1`,
		where,
	)
}

func userQuery(paramTyp int) string {
	var where string
	switch paramTyp {
	case typID:
		where = "id"
	case typUser:
		where = "lower(u.username)"
	}

	return fmt.Sprintf(
		`select u.id, u.password, u.username, u.fullname 
		from users u
		where u.is_active = true and %s = $1
		limit 1`,
		where,
	)
}

// GetUserByIDQuery ...
func (pr *Repository) GetUserByIDQuery(userID int) ([][]string, error) {
	qapi := userQuery(typID)
	return pr.Connection.ExecuteDBQuery(qapi, userID)
}

// GetUserByUsernameQuery ...
func (pr *Repository) GetUserByUsernameQuery(username string) ([][]string, error) {
	qapi := userQuery(typUser)
	return pr.Connection.ExecuteDBQuery(qapi, utils.TrimLower(username))
}

// GetUserDataByIDQuery ...
func (pr *Repository) GetUserDataByIDQuery(userID int) ([][]string, error) {
	qapi := userInfoQuery(typID)
	return pr.Connection.ExecuteDBQuery(qapi, userID)
}

// GetUserDataByUsernameQuery ...
func (pr *Repository) GetUserDataByUsernameQuery(username string) ([][]string, error) {
	qapi := userInfoQuery(typUser)
	return pr.Connection.ExecuteDBQuery(qapi, utils.TrimLower(username))
}

// GetUserPermissionQuery ...
func (pr *Repository) GetUserPermissionQuery(menu string) ([][]string, error) {
	qapi := "select lower(menu) as menu, description from user_permission where lower(menu) = $1"
	return pr.Connection.ExecuteDBQuery(qapi, utils.TrimLower(menu))
}

// GetUserPermissionBulkQuery ...
func (pr *Repository) GetUserPermissionBulkQuery(menu []string) ([][]string, error) {
	var (
		menus []string
		args  []interface{}
	)

	for i, m := range menu {
		menus = append(menus, "$"+strconv.Itoa(i+1))
		args = append(args, utils.TrimLower(m))
	}

	qapi := fmt.Sprintf(
		"select lower(menu) as menu, description from user_permission where lower(menu) in (%s)",
		strings.Join(menus, ","),
	)

	return pr.Connection.ExecuteDBQuery(qapi, args...)
}

// GetUserRoleQuery ...
func (pr *Repository) GetUserRoleQuery(name string) ([][]string, error) {
	qapi := `select lower(name) as name, description, lower(permissions::varchar) as permission 
	from user_role where lower(name) = $1`
	return pr.Connection.ExecuteDBQuery(qapi, utils.TrimLower(name))
}
