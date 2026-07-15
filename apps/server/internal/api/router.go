package api

import (
	"encoding/json"
	"net/http"

	"allone/server/internal/app"
	"allone/server/internal/auth"
	"allone/server/internal/device"
	"allone/server/internal/middleware"
	"allone/server/internal/websocket"
	"github.com/go-chi/chi/v5"
)

func NewRouter(a *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(a.Logger))
	r.Get("/health", HealthHandler(a))

	// 1. Declare these once at the top of the router setup
	jwtService := auth.NewJWTService(a.Config.JWTSecret)
	deviceRepo := device.NewRepository(a.DB)

	// Websocket setup
	wsHandler := websocket.NewHandler(
		a.Hub,
		jwtService,
		deviceRepo,
	)
	r.Get("/ws", wsHandler.Connect)

	// Auth Initialization 
	repo := auth.NewRepository(a.DB)
	// FIX: Removed the duplicate 'jwtService :=' declaration line entirely 
	// since we already initialized it above.
	service := auth.NewService(repo, jwtService)
	handler := auth.NewHandler(service)

	// Device Initialization
	// FIX: Removed the duplicate 'deviceRepo :=' declaration line entirely
	// since we already initialized it above.
	deviceService := device.NewService(deviceRepo)
	deviceHandler := device.NewHandler(deviceService)

	// Public Auth Routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	})

	// Protected Routes (Require JWT Authentication)
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware(jwtService))

		// User Profile Route
		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			id := r.Context().Value(auth.UserIDKey)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"user_id": id,
			})
		})

		// Device Routes
		r.Route("/devices", func(r chi.Router) {
			r.Post("/register", deviceHandler.Register)
			r.Get("/", deviceHandler.List)
		})
	})

	return r
}