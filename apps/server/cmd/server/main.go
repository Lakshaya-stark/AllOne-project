package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log/slog"

	"allone/server/internal/api"
	"allone/server/internal/config"
	"allone/server/internal/database"
	"allone/server/internal/app"
	"allone/server/internal/logger"
)

func main() {

	cfg := config.Load()
	
	log := logger.New()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Error("Failed to connect PostgreSQL",
		slog.Any("error", err),
		)
		os.Exit(1)
	}
	defer db.Close()
	log.Info("Connected to PostgreSQL")
	
	
	redisClient, err := database.NewRedis(cfg)
	if err != nil {
		log.Error("Failed to connect PostgreSQL",
		slog.Any("error", err),
		)

		os.Exit(1)
	}
	defer redisClient.Close()
		application := &app.App{
		Config: cfg,
		DB:     db,
		Redis: redisClient,
		Logger: log,
	}
	log.Info("Connected to Redis")
	
	
	router := api.NewRouter(application)
	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {

		fmt.Printf("Server running on port %s\n", cfg.AppPort)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			panic(err)
		}
	}()

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

	server.Shutdown(ctx)

	fmt.Println("Server stopped")
}