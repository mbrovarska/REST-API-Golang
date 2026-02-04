package main

import (
	"example.com/rest-api-notes/internal/config"
	"example.com/rest-api-notes/internal/routes"
	httpserver "example.com/rest-api-notes/pkg/http-server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(config.Module, fx.Provide(
		zap.NewDevelopment,
        httpserver.NewServer,
	),
		routes.Module,
		fx.Invoke(func(log *zap.Logger) {
			log.Info("Application initialized and starting...")

		}),
	).Run()
	
}