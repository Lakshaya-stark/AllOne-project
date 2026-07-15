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

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(auth.UserIDKey).(uuid.UUID)

	err = h.service.Register(
		r.Context(),
		userID,
		req,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(RegisterResponse{
		Message: "Device Registered Successfully",
	})
}

func (h *Handler) List(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().Value(auth.UserIDKey).(uuid.UUID)

	devices, err := h.service.ListDevices(
		r.Context(),
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(devices)
}