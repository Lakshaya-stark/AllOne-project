package api

import (
	"encoding/json"
	"net/http"

	"allone/server/internal/app"
	"allone/server/internal/auth"
	"allone/server/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(a *app.App) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(a.Logger))
	r.Get("/health", HealthHandler(a))

	repo := auth.NewRepository(a.DB)
	jwtService := auth.NewJWTService(a.Config.JWTSecret)
	service := auth.NewService(repo, jwtService)
	handler := auth.NewHandler(service)

	// Public Auth Routes
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	})

	// Protected Routes (Require JWT Authentication)
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware(jwtService))

		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			id := r.Context().Value(auth.UserIDKey)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"user_id": id,
			})
		})
	})

	return r
}

