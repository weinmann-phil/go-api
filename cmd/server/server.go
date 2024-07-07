package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weinmann-phil/gobank/_common/interfaces"
)

type GinServer interface {
	Start(ctx context.Context, httpAddress string) error
	Shutdown(ctx context.Context) error
	RegisterGroupRoute(path string, routes []interfaces.RouteDefinition, middleware ...gin.HandlerFunc)
	RegisterRouter(method, path string, handler gin.HandlerFunc)
}

type GinServerBuilder struct {
}

type ginServer struct {
	engine *gin.Engine
	server *http.Server
}

func NewGinServerBuilder() *GinServerBuilder {
	return &GinServerBuilder{}
}

// Build a GinServer
func (builder *GinServerBuilder) Build() GinServer {
	engine := gin.Default()
	return &ginServer{engine: engine}
}

// Start up HTTP server.
func (gs *ginServer) Start(ctx context.Context, httpAddress string) error {
	gs.server = &http.Server{
		Addr:    httpAddress,
		Handler: gs.engine,
	}

	go func() {
		if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Listening %s \n", err)
		}
	}()

	logrus.Infof("HTTP server running on port %s", httpAddress)
	return nil
}

// Shutdown server gracefully.
func (gs *ginServer) Shutdown(ctx context.Context) error {
	if err := gs.server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shutdown %s", err)
	}

	logrus.Infof("Server is shutting down.")

	return nil
}

// Method to register a single route
func (gs *ginServer) RegisterRouter(method, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		gs.engine.GET(path, handler)
	case "POST":
		gs.engine.POST(path, handler)
	case "PUT":
		gs.engine.PUT(path, handler)
	case "DELETE":
		gs.engine.DELETE(path, handler)
	case "PATCH":
		gs.engine.PATCH(path, handler)
	default:
		logrus.Errorf("Invalid http method")
	}
}

// Method to register a group route
func (gs *ginServer) RegisterGroupRoute(path string, routes []interfaces.RouteDefinition, middleWare ...gin.HandlerFunc) {
	group := gs.engine.Group(path)
	group.Use(middleWare...)

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Path, route.Handler)
		case "POST":
			group.POST(route.Path, route.Handler)
		case "PUT":
			group.PUT(route.Path, route.Handler)
		case "DELETE":
			group.DELETE(route.Path, route.Handler)
		case "PATCH":
			group.PATCH(route.Path, route.Handler)
		default:
			logrus.Errorf("Invalid http method")
		}
	}
}
