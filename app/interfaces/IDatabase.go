package interfaces

import "gorm.io/gorm"

type IDatabase interface {
	GetInstance() *gorm.DB
}
