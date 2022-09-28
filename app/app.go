package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shasw94/projX/app/api"
	"github.com/shasw94/projX/app/dbs"
	"github.com/shasw94/projX/app/repositories"
	"github.com/shasw94/projX/app/router"
	"github.com/shasw94/projX/app/services"
	"github.com/shasw94/projX/logger"
	"github.com/shasw94/projX/pkg/jwt"
	"go.uber.org/dig"
	"time"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	auth, err := InitAuth()
	_ = container.Provide(func() jwt.IJWTAuth {
		return auth
	})

	// Inject database
	err = dbs.Inject(container)
	if err != nil {
		logger.Error("Failed to inject database", err)
	}

	err = repositories.Inject(container)
	if err != nil {
		logger.Error("Failed to inject repositories", err)
	}

	err = services.Inject(container)
	if err != nil {
		logger.Error("Failed to inject services", err)
	}

	err = api.Inject(container)
	if err != nil {
		logger.Error("Failed to inject APIs", err)
	}

	return container
}

// InitGinEngine initial new gin engine
func InitGinEngine(container *dig.Container) *gin.Engine {
	app := gin.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "http://localhost:3000"
		//},
		MaxAge: 12 * time.Hour,
	}))
	router.Docs(app)
	err := router.RegisterAPI(app, container)
	if err != nil {
		return nil
	}

	return app
}
