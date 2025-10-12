package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Config struct {
	Port int
	Host string
	Pass string
	Db   int
}

// Init initiate Redis library
func Init(cfg *Config, l *zerolog.Logger) (*redis.Client, error) {
	var (
		log = l.With().Logger()
	)
	if cfg.Port == 0 {
		cfg.Port = 6379
	}
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Password: cfg.Pass,
		DB:       cfg.Db,
	})

	status := rdb.Ping(context.Background())
	if status.Err() != nil {
		return nil, status.Err()
	}

	log.Info().Msgf("Redis client initialized on %v/%v", rdb.Options().Addr, rdb.Options().DB)

	return rdb, nil
}
