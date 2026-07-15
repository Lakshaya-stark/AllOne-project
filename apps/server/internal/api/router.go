package api

import (
	"encoding/json"
	"net/http"

	"allone/server/internal/app"
	"allone/server/internal/auth"
	"allone/server/internal/device" 
	"allone/server/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(a *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(a.Logger))
	r.Get("/health", HealthHandler(a))

	// Auth Initialization
	repo := auth.NewRepository(a.DB)
	jwtService := auth.NewJWTService(a.Config.JWTSecret)
	service := auth.NewService(repo, jwtService)
	handler := auth.NewHandler(service)

	// Device Initialization
	deviceRepo := device.NewRepository(a.DB)
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