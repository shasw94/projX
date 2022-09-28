package pivot

// UserRole represents the database model of user roles relationships
type UserRole struct {
	UserID string `gorm:"primaryKey" json:"user_id"`
	RoleID string `gorm:"primaryKey" json:"role_id"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
