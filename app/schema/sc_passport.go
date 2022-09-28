package schema

import "github.com/shasw94/projX/pkg/utils"

type RoleOption struct {
	WithPermissions bool
	Pagination      *utils.Pagination
}

type PermissionOption struct {
	Pagination *utils.Pagination
}
