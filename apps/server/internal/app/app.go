package app

import (
	"log/slog"

	"allone/server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Config *config.Config

	DB *pgxpool.Pool

	Redis *redis.Client

	Logger *slog.Logger
}