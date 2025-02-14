package cmd

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/weinmann-phil/go-api/cmd/server"
	"github.com/weinmann-phil/go-api/config"
	"github.com/weinmann-phil/go-api/provider"
)

func Execute() {
	builder := server.NewGinServerBuilder()
	server := builder.Build()

	ctx := context.Background()
	logrus.Info("0")
	// config.LoadEnvironment()

	db, err := config.SetupDatabase()
	if err != nil {
		logrus.Fatalf("Error setting up database %v", err)
	}

	provider.NewProvider(db, server)
	go func() {
		if err := server.Start(ctx, os.Getenv(config.AppPort)); err != nil {
			logrus.Errorf("Error starting server %v", err)
		}
	}()

	<-ctx.Done()
	logrus.Info("Server stopped")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("Error stopping the server %v", err)
	}
}
