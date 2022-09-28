package models

import (
	"github.com/shasw94/projX/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	Model        `json:"inline"`
	Username     string `json:"username" gorm:"unique;not null;index"`
	Email        string `json:"email" gorm:"unique;not null;index"`
	Password     string `json:"password" gorm:"not null;index"`
	Role         Role
	RoleID       string `json:"role_id" gorm:"not null;index"`
	RefreshToken string `json:"refresh_token" gorm:"size:500;index"`
	FullName     string `json:"full_name"`
	ProfileImage string `json:"profile_image"`
	Mobile       string `json:"mobile" gorm:"not null;default:0"`
}

// BeforeCreate handle before create user
func (u *User) BeforeCreate(db *gorm.DB) error {
	err := u.Model.BeforeCreate(db)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword([]byte(u.Password))
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	return nil
}
