package db

import (
	"os"
)

type DBConfig struct {
	ConnectionString string
}


func NewDBConfig() DBConfig {
	connString := os.Getenv("DB_URL")
	if connString == ""{
		panic("DATABASE_URL environment variable is required")
	}

	return DBConfig{
		ConnectionString: connString,
	}
}