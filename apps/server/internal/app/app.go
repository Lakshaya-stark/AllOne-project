package app

import (
	"log/slog"

	"allone/server/internal/config"
	"allone/server/internal/websocket"
	"allone/server/internal/presence"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Config *config.Config

	DB *pgxpool.Pool
	Hub *websocket.Hub

Presence *presence.Service
	Redis *redis.Client

	Logger *slog.Logger
}

