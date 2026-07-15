package device

import (
	"encoding/json"
	"net/http"

	"allone/server/internal/auth"

	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(auth.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "invalid user context", http.StatusUnauthorized)
		return
	}

	device, err := h.service.Register(
		r.Context(),
		userID,
		req,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(RegisterResponse{
		Message:  "Device registered successfully",
		DeviceID: device.ID.String(),
	})
}

func (h *Handler) List(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID, ok := r.Context().Value(auth.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "invalid user context", http.StatusUnauthorized)
		return
	}

	devices, err := h.service.ListDevices(
		r.Context(),
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(devices)
}