package models

type Story struct {
	Model       `json:"inline"`
	Name        string `json:"name" gorm:"unique;not null;index"`
	Description string `json:"description" gorm:"not null;"`
	CoverImage  string `json:"cover_image" gorm:"not null;"`
}
