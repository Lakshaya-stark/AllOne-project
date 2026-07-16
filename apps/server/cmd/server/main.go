package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"allone/server/internal/presence"
	"allone/server/internal/api"
	"allone/server/internal/app"
	"allone/server/internal/config"
	"allone/server/internal/database"
	"allone/server/internal/logger"
	"allone/server/internal/websocket"
)

func main() {

	// ----------------------------
	// Load Configuration
	// ----------------------------
	cfg := config.Load()

	// ----------------------------
	// Initialize Logger
	// ----------------------------
	log := logger.New()

	// ----------------------------
	// PostgreSQL
	// ----------------------------
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Error("Failed to connect PostgreSQL",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
	defer db.Close()

	log.Info("Connected to PostgreSQL")

	// ----------------------------
	// Redis
	// ----------------------------
	redisClient, err := database.NewRedis(cfg)
	if err != nil {
		log.Error("Failed to connect Redis",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
	defer redisClient.Close()

	log.Info("Connected to Redis")

	redisPresence := presence.NewRedisRepository(redisClient)

postgresPresence := presence.NewPostgresRepository(db)

presenceService := presence.NewService(
	redisPresence,
	postgresPresence,
)

	// ----------------------------
	// WebSocket Hub
	// ----------------------------
	hub := websocket.NewHub()

	go hub.Run()

	log.Info("WebSocket Hub started")

	// ----------------------------
	// Dependency Container
	// ----------------------------
	application := &app.App{
	Config: cfg,
	DB: db,
	Redis: redisClient,
	Logger: log,
	Hub: hub,

	Presence: presenceService,
}

	// ----------------------------
	// HTTP Router
	// ----------------------------
	router := api.NewRouter(application)

	server := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// ----------------------------
	// Start Server
	// ----------------------------
	go func() {

		log.Info(
			"HTTP Server Started",
			slog.String("port", cfg.AppPort),
		)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Error(
				"HTTP Server Error",
				slog.Any("error", err),
			)

			os.Exit(1)
		}
	}()

	// ----------------------------
	// Graceful Shutdown
	// ----------------------------
	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-stop

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {

		log.Error(
			"Graceful shutdown failed",
			slog.Any("error", err),
		)
	}

	log.Info("Server stopped")
}