package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Config struct {
	Port    int `yaml:"port"`
	Version string
	Env     string
}


func NewServer(lc fx.Lifecycle, cfg Config) *gin.Engine {

	router := gin.Default()
	
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Printf("Starting server on port %d (Version: %s, Env: %s)", cfg.Port, cfg.Version, cfg.Env)
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Server failed: %v", err)
				}
			}()
			return  nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return srv.Shutdown(ctx)
		},
	})

	return router


}