package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primaryKey;"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;index"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate handle before create
func (model *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if model.ID == "" {
		model.ID = uuid.New().String()
	}
	if model.CreatedAt.IsZero() {
		model.CreatedAt = time.Now()
	}
	if model.UpdatedAt.IsZero() {
		model.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate handle before update
func (model *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	model.UpdatedAt = time.Now().UTC()
	return nil
}
