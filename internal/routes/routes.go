package routes

import (
	"example.com/rest-api-notes/internal/handler"
	"example.com/rest-api-notes/internal/health"
	"example.com/rest-api-notes/internal/repository"
	"example.com/rest-api-notes/internal/service"
	httpserver "example.com/rest-api-notes/pkg/http-server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

var Module = fx.Module("routes",
	fx.Provide(
		repository.NewUserRepository,
		service.NewUserService,
		handler.NewUserHandler,
	),
	fx.Invoke(RegisterRoutes),
)

func RegisterRoutes(
	router *gin.Engine, 
	config httpserver.Config,
	userHandler *handler.UserHandler,
	) {
	// API routes
	api := router.Group("/api")
	{
		api.GET("/health", health.CheckHealth(config.Version))
		auth := api.Group("auth")
		{
			auth.POST("/signup", userHandler.SignUp)
		}
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}