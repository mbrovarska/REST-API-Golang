package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("database",
	fx.Provide(
		NewDBConfig,
		NewPool,
	),
)

func NewPool(lc fx.Lifecycle, dbConfig DBConfig, logger *zap.Logger) (*pgxpool.Pool, error) {
	logger.Info("Initializing database connection pool")
	config, err := pgxpool.ParseConfig(dbConfig.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse supabase config: %w", err)
	}

	//manage connections
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30

	// disable prepared statement cache to avoid conflicts with Squirrel
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

    ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to supabase: %w", err)
	}

	//verify connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("supabase ping failed: %w", err)
	}
	logger.Info("Database connection pool initialized successfully",
		zap.Int32("max_connections", config.MaxConns),
		zap.Int32("min_connections", config.MinConns),
	)

	//run migrations
	if err := RunMigrations(pool, logger); err != nil {
		pool.Close()
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	//register lifecycle hooks
	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			logger.Info("Closing database connection pool")
			pool.Close()
			return nil
		},
	})


	return pool, nil

}