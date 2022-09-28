package pivot

type UserPermission struct {
	UserID       string `gorm:"primaryKey" json:"user_id"`
	PermissionID string `gorm:"primaryKey" json:"permission_id"`
}

func (UserPermission) TableName() string {
	return "user_permissions"
}
