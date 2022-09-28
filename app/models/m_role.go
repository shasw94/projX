package models

type Role struct {
	Model       `json:"inline"`
	Name        string `json:"name" gorm:"unique;not null;index"`
	GuardName   string `gorm:"size:255;not null;index" json:"guard_name"`
	Description string `json:"description" gorm:"size:255;"`

	// Many to Many
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"permissions"`
}
