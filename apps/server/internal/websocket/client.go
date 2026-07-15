package websocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID   uuid.UUID
	DeviceID uuid.UUID

	Conn *websocket.Conn

	Send chan []byte

	Hub *Hub
}