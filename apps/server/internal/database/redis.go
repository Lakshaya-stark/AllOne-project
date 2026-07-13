package database

import (
	"context"
	"fmt"

	"allone/server/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *config.Config) (*redis.Client, error) {

	addr := fmt.Sprintf("%s:%s",
		cfg.RedisHost,
		cfg.RedisPort,
	)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: cfg.RedisPassword,
		DB: 0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}