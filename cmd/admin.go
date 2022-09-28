package main

import (
	"github.com/shasw94/projX/app"
	"github.com/shasw94/projX/app/migration"
	"github.com/shasw94/projX/logger"
)

func main() {
	container := app.BuildContainer()
	err := migration.CreateAdmin(container)
	if err != nil {
		logger.Error("Failed to create admin: ", err)
	}
}
