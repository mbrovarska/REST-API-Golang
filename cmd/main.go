package main

import (
	"log"

	"example.com/rest-api-notes/internal/config"
	"example.com/rest-api-notes/internal/db"
	"example.com/rest-api-notes/internal/routes"
	httpserver "example.com/rest-api-notes/pkg/http-server"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	//load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	fx.New(
		config.Module, 
		db.Module,
		fx.Provide(
			zap.NewDevelopment,
			httpserver.NewServer,
	),
		routes.Module,
		fx.Invoke(func(log *zap.Logger, pool *pgxpool.Pool) {
			log.Info("Application initialized and starting...")
			log.Info("Database pool ready", 
				zap.Int("max_conns", int(pool.Config().MaxConns)))

		}),
	).Run()
	
}