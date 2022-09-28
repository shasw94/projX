package migration

import (
	"fmt"
	"github.com/shasw94/projX/app/interfaces"
	"github.com/shasw94/projX/app/models"
	"go.uber.org/dig"
)

// CreateAdmin create new user role admin
func CreateAdmin(container *dig.Container) error {
	return container.Invoke(func(
		userRepo interfaces.IUserRepository,
		roleRepo interfaces.IRoleRepository,
	) error {
		fmt.Println("reached here")
		adminRole := &models.Role{Name: "admin", Description: "Admin"}
		userRole := &models.Role{Name: "user", Description: "User"}
		err := roleRepo.Create(adminRole)
		err = roleRepo.Create(userRole)
		if err != nil {
			return err
		}
		err = userRepo.Create(&models.User{
			Username: "admin",
			Password: "admin",
			Email:    "admin@admin.com",
			RoleID:   adminRole.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

// Migrate migrate to database
func Migrate(container *dig.Container) error {
	return container.Invoke(func(db interfaces.IDatabase) error {
		User := models.User{}
		Role := models.Role{}

		db.GetInstance().AutoMigrate(&User, &Role)
		return nil
	})
}
