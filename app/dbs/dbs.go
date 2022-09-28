package dbs

import (
	"fmt"
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/config"
	"github.com/shasw94/projX/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

// NewDatabase returns new IDatabase interface
func NewDatabase() interfaces.IDatabase {
	dbConfig := config.Config.Database
	connectionPath := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	logger.Info(connectionPath)

	db, err := gorm.Open(mysql.Open(connectionPath), &gorm.Config{})
	if err != nil {
		logger.Fatal("Cannot connect to database ", err)
	}

	sqlDB, err := db.DB()
	// Set up connection pool
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)

	return &database{
		db: db,
	}
}

func (d *database) GetInstance() *gorm.DB {
	return d.db
}

