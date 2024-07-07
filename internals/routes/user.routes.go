package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weinmann-phil/gobank/_common/interfaces"
	"github.com/weinmann-phil/gobank/cmd/server"
	"github.com/weinmann-phil/gobank/internals/handler"
)

func RegisterUserRoutes(server server.GinServer, userHandler *handler.UserHandler) {
	server.RegisterGroupRoute("/api/v1/user", []interfaces.RouteDefinition{
		{Method: "POST", Path: "/register", Handler: userHandler.CreateUser},
	}, func(c *gin.Context) {
		logrus.Infof("Request on %s", c.Request.URL.Path)
	})
}

// func GetUserRoutes(server server.GinServer, userHandler *handler.UserHandler) {
// 	server.RegisterGroupRoute(("/api/v1/user", []inter[]interfaces.RouteDefinition{
// 		{Method: "GET", Path: "/", Handler: userHandler}
// 	}))
// }
