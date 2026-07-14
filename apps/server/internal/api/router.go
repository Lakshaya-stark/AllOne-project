package api

import (
	"allone/server/internal/app"
	"allone/server/internal/middleware"
	"allone/server/internal/auth"

	"github.com/go-chi/chi/v5"
)

func NewRouter(a *app.App) *chi.Mux {
	
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger(a.Logger))
	r.Get("/health", HealthHandler(a))

	repo := auth.NewRepository(a.DB)
service := auth.NewService(repo)
handler := auth.NewHandler(service)

r.Route("/auth", func(r chi.Router) {
	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)
})

	return r
}