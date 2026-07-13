package api

import (
	"allone/server/internal/app"
	"allone/server/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(a *app.App) *chi.Mux {
	
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(a.Logger))
	r.Get("/health", HealthHandler(a))

	return r
}