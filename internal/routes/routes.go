package routes

import (
	"example.com/rest-api-notes/internal/health"
	httpserver "example.com/rest-api-notes/pkg/http-server"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("routes", 
	fx.Invoke(RegisterRoutes),
)

func RegisterRoutes(r *gin.Engine, cfg httpserver.Config) {
	api := r.Group("/api")
	{
		api.GET("/health", gin.WrapF(health.CheckHealth(cfg.Version)))
	}
}