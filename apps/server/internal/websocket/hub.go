package websocket

import (
	"context"
	"sync"

	"allone/server/internal/presence"

	"github.com/google/uuid"
)

type Hub struct {
	// Keep these lowercase (unexported) to avoid collision with Register/Unregister methods
	clients    map[uuid.UUID]*Client
	register   chan *Client
	unregister chan *Client

	Presence *presence.Service
	mu       sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uuid.UUID]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.DeviceID] = client
			h.mu.Unlock()

		case client := <-h.unregister: // Changed to lowercase h.unregister
			if h.Presence != nil {
				_ = h.Presence.SetOffline(
					context.Background(),
					client.DeviceID,
				)
			}

			h.mu.Lock()
			if existing, ok := h.clients[client.DeviceID]; ok { // Changed to lowercase h.clients
				if existing == client {
					delete(h.clients, client.DeviceID) // Changed to lowercase h.clients
					close(client.Send)
				}
			}
			h.mu.Unlock()
		}
	}
}

// Exported method to register a client (sends to the unexported register channel)
func (h *Hub) Register(client *Client) {
	h.register <- client
}

// Exported method to unregister a client (sends to the unexported unregister channel)
func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

func (h *Hub) Client(deviceID uuid.UUID) (*Client, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	client, ok := h.clients[deviceID] // Correctly uses lowercase h.clients
	return client, ok
}