package middleware

import (
	"context"
	"fmt"
	"strings"

	"go-microservices.org/api/graph/model"
	servicemodel "go-microservices.org/core/proto"
	"go-microservices.org/core/utils"

	"github.com/99designs/gqlgen/graphql"
)

const (
	accessDeniedPermission = "access denied to this menu"
)

// PermissionCheck ...
func PermissionCheck(
	ctx context.Context, obj interface{}, next graphql.Resolver, control []model.AccessControl,
) (interface{}, error) {
	if !isHasPermission(ctx, control) {
		return nil, fmt.Errorf(accessDeniedPermission)
	}

	return next(ctx)
}

func isHasPermission(ctx context.Context, control []model.AccessControl) bool {
	authData := GetTokenContext(ctx)
	if authData == nil {
		return false
	}

	fc := graphql.GetFieldContext(ctx)

	var accessControl []string
	if accessControl = getAccessControl(fc.Path().String(), authData.ParsedToken.Permissions); accessControl == nil {
		return false
	}

	for _, ctr := range control {
		if utils.FindString(accessControl, strings.ToLower(ctr.String())) {
			return true
		}
	}

	return false
}

func getAccessControl(currentPath string, permissions []*servicemodel.AccessPermission) []string {
	var result []string
	currentPath = strings.ToLower(currentPath)

	for _, perm := range permissions {
		if perm.Menu == "*" || perm.Menu == currentPath {
			result = perm.Control
			break
		}
	}

	return result
}
