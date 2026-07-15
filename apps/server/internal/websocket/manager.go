package websocket

import "github.com/google/uuid"

type Manager struct {

	hub *Hub
}

func NewManager(

	hub *Hub,

) *Manager {

	return &Manager{

		hub: hub,
	}
}


func (m *Manager) Send(

	deviceID uuid.UUID,

	data []byte,

) bool {

	client, ok := m.hub.Client(deviceID)

	if !ok {

		return false
	}

	client.Send <- data

	return true
}

func (m *Manager) Broadcast(

	data []byte,

) {

	m.hub.mu.RLock()

	defer m.hub.mu.RUnlock()

	for _, client := range m.hub.clients {

		client.Send <- data
	}
}