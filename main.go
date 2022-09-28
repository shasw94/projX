package main

import (
	"context"
	"github.com/shasw94/projX/app"
	"github.com/shasw94/projX/app/migration"
	_ "github.com/shasw94/projX/docs"
	"github.com/shasw94/projX/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ProductionEnv = "production"
)

func main() {
	logger.Initialize(ProductionEnv)
	container := app.BuildContainer()
	engine := app.InitGinEngine(container)

	err := migration.Migrate(container)
	if err != nil {
		logger.Warn("Failed to migrate data: ", err)
	}

	server := &http.Server{
		Addr:    ":8888",
		Handler: engine,
	}

	go func() {
		logger.Info("Listen at:", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 1 seconds
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be a catch, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}
	// catching ctx.Done(). timeout of 1 secs
	select {
	case <-ctx.Done():
		logger.Info("Timeout of 1 seconds")
	}
	logger.Info("server exiting")
}
