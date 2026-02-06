package routes

import (
	"example.com/rest-api-notes/internal/health"
	httpserver "example.com/rest-api-notes/pkg/http-server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

var Module = fx.Module("routes",
	fx.Invoke(RegisterRoutes),
)

func RegisterRoutes(router *gin.Engine, config httpserver.Config) {
	// API routes
	api := router.Group("/api")
	{
		api.GET("/health", health.CheckHealth(config.Version))
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}