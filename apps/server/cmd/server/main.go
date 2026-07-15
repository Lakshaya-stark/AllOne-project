package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"allone/server/internal/websocket"

	"allone/server/internal/api"
	"allone/server/internal/app"
	"allone/server/internal/config"
	"allone/server/internal/database"
	"allone/server/internal/logger"
	
)

func main() {

	cfg := config.Load()

	log := logger.New()

	// 1. Initialize PostgreSQL
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Error("Failed to connect PostgreSQL",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
	defer db.Close()
	log.Info("Connected to PostgreSQL")

	// 2. Initialize Redis
	redisClient, err := database.NewRedis(cfg)
	if err != nil {
		log.Error("Failed to connect Redis", 
			slog.Any("error", err),
		)
		os.Exit(1)
	}
	defer redisClient.Close()
	log.Info("Connected to Redis")

	// 3. Initialize WebSocket Hub
	hub := websocket.NewHub() 
	go hub.Run()
	log.Info("WebSocket Hub is running")

	// 4. Build Application Dependency Container
	application := &app.App{
		Config: cfg,
		DB:     db,
		Redis:  redisClient,
		Logger: log,
		Hub:    hub,
	}

	// 5. Setup Router and Server
	router := api.NewRouter(application)
	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 6. Start Server HTTP Loop
	go func() {
		fmt.Printf("Server running on port %s\n", cfg.AppPort)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// 7. Graceful Shutdown Listeners
	stop := make(chan os.Signal, 1)
	signal.Notify(stop,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-stop

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", slog.Any("error", err))
	}

	fmt.Println("Server stopped")
}