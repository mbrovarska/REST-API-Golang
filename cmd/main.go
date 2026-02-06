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

// @title           Notes API
// @version         1.0
// @description     A REST API for managing notes with user authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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