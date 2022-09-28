package models

type Permission struct {
	Model
	Name        string `gorm:"size:255;not null" json:"name"`
	GuardName   string `gorm:"size:255;not null" json:"guard_name"`
	Description string `gorm:"size:255;not null;index" json:"description"`
}
