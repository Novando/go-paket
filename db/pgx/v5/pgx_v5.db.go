package pgx

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Config struct {
	User    string
	Pass    string
	Host    string
	Port    uint
	Name    string
	Schema  string
	MaxPool uint
	SSL     bool
}

// InitPGXv5 initialize database connection
func InitPGXv5(cfg Config, l *zerolog.Logger) (*pgxpool.Pool, error) {
	var (
		log = l.With().Logger()
	)

	if cfg.Name == "" {
		return nil, errors.New("DB_NAME_REQUIRED")
	}
	if cfg.MaxPool == 0 {
		cfg.MaxPool = 5
	}
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	if cfg.Port == 0 {
		cfg.Port = 5432
	}
	if cfg.User == "" {
		cfg.User = "postgres"
	}
	if cfg.Schema == "" {
		cfg.Schema = "public"
	}
	sslMode := "disable"
	if cfg.SSL {
		sslMode = "require"
	}
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?pool_max_conns=%d&search_path=%s&sslmode=%s",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.MaxPool,
		cfg.Schema,
		sslMode,
	)
	c, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}
	if err = pool.Ping(context.Background()); err != nil {
		log.Err(err).Send()
		return nil, err
	}
	log.Info().Msgf("PG connected with %s@%s", cfg.Name, cfg.Host)
	return pool, err
}
