package websocket

import "sync"

type Hub struct {

	Clients map[string]*Client

	Register chan *Client

	Unregister chan *Client

	mu sync.RWMutex
}

func NewHub() *Hub {

	return &Hub{

		Clients: make(map[string]*Client),

		Register: make(chan *Client),

		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {

	for {

		select {

		case client := <-h.Register:

			h.mu.Lock()

			h.Clients[client.DeviceID.String()] = client

			h.mu.Unlock()

		case client := <-h.Unregister:

			h.mu.Lock()

			delete(h.Clients, client.DeviceID.String())

			close(client.Send)

			h.mu.Unlock()
		}
	}
}