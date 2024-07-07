package provider

import (
	"github.com/weinmann-phil/gobank/cmd/server"
	"github.com/weinmann-phil/gobank/internals/handler"
	"github.com/weinmann-phil/gobank/internals/repository"
	"github.com/weinmann-phil/gobank/internals/routes"
	"github.com/weinmann-phil/gobank/internals/services"
	"gorm.io/gorm"
)

func NewProvider(db *gorm.DB, server server.GinServer) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	routes.RegisterUserRoutes(server, userHandler)
}
