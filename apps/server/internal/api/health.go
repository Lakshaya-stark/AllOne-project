package api

import (
	"encoding/json"
	"net/http"
	"allone/server/internal/app"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	}
}