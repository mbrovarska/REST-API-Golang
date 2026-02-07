package config

import (
	"os"

	"go.uber.org/fx"
)

type AuthConfig struct {
	JWTSecret string
}

var Module = fx.Module("config", fx.Provide(
	NewConfig,
	NewAuthConfig,
 ),
)

func NewAuthConfig() AuthConfig {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == ""{
		jwtSecret = "dev-secret-change-this-in-production"
	}

	return AuthConfig{
		JWTSecret: jwtSecret,
	}
}