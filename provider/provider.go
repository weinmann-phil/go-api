package provider

import (
	"github.com/weinmann-phil/go-api/cmd/server"
	"github.com/weinmann-phil/go-api/internals/handler"
	"github.com/weinmann-phil/go-api/internals/repository"
	"github.com/weinmann-phil/go-api/internals/routes"
	"github.com/weinmann-phil/go-api/internals/services"
	"gorm.io/gorm"
)

func NewProvider(db *gorm.DB, server server.GinServer) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	routes.RegisterUserRoutes(server, userHandler)
}
