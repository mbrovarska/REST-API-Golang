package config

import (
	"os"
	"strconv"

	httpserver "example.com/rest-api-notes/pkg/http-server"
)


	
	
func NewConfig() httpserver.Config {
	//port configaration
	portStr := os.Getenv("HTTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8080
	}

	//environment configuration
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	//handle version 
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "1.0.0"
	}

	return httpserver.Config{
		Port: port,
		Env: env,
		Version: version,
	}
 }
	
