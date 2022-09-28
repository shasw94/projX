package schema

import (
	"github.com/shasw94/projX/app/models"
	"github.com/shasw94/projX/pkg/utils"
)

type Roles []models.Role

func (s Roles) Origin() []models.Role {
	return []models.Role(s)
}

func (s Roles) Len() (length int64) {
	return int64(len(s))
}

func (s Roles) GuardNames() (guards []string) {
	for _, role := range s {
		guards = append(guards, role.GuardName)
	}
	return guards
}

func (s Roles) IDs() (IDs []string) {
	for _, role := range s {
		IDs = append(IDs, role.ID)
	}
	return IDs
}

func (s Roles) Permissions() (permissions Permission) {
	var IDs []string
	for _, a := range s {
		if len(a.Permissions) > 0 {
			for _, prm := range a.Permissions {
				if !utils.InArray(prm.ID, IDs) {
					permissions = append(permissions, prm)
				}
				IDs = append(IDs, prm.ID)
			}
		}
	}
	return
}
