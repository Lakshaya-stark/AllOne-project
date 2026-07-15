package websocket

import (
	"net/http"
	"strings"

	"allone/server/internal/auth"
	"allone/server/internal/device"

	"github.com/google/uuid"
	gws "github.com/gorilla/websocket"
)

var upgrader = gws.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handler struct {
	Hub        *Hub
	JWT        *auth.JWTService
	DeviceRepo device.Repository
}

func NewHandler(
	hub *Hub,
	jwt *auth.JWTService,
	repo device.Repository,
) *Handler {

	return &Handler{
		Hub:        hub,
		JWT:        jwt,
		DeviceRepo: repo,
	}
}

func (h *Handler) Connect(
	w http.ResponseWriter,
	r *http.Request,
) {

	// -----------------------------
	// Authorization Header
	// -----------------------------
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		http.Error(w, "missing authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := h.JWT.Validate(tokenString)

	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	// -----------------------------
	// Device ID
	// -----------------------------
	deviceIDString := r.URL.Query().Get("device_id")

	if deviceIDString == "" {
		http.Error(w, "missing device_id", http.StatusBadRequest)
		return
	}

	deviceID, err := uuid.Parse(deviceIDString)

	if err != nil {
		http.Error(w, "invalid device_id", http.StatusBadRequest)
		return
	}

	// -----------------------------
	// Verify Device
	// -----------------------------
	dev, err := h.DeviceRepo.GetByID(
		r.Context(),
		deviceID,
	)

	if err != nil {
		http.Error(w, "device not found", http.StatusUnauthorized)
		return
	}

	if dev.UserID != claims.UserID {
		http.Error(w, "device does not belong to user", http.StatusUnauthorized)
		return
	}

	// -----------------------------
	// Upgrade Connection
	// -----------------------------
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	// -----------------------------
	// Create Client
	// -----------------------------
	client := NewClient(
		conn,
		claims.UserID,
		deviceID,
		h.Hub,
	)

	h.Hub.Register(client)

	go client.ReadPump()
	go client.WritePump()
}