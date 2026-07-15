package websocket

import (
	"time"

	"github.com/google/uuid"
	gws "github.com/gorilla/websocket"
)

type Client struct {

	ID ClientID

	UserID uuid.UUID

	DeviceID uuid.UUID

	Conn *gws.Conn

	Send chan []byte

	Hub *Hub

	LastSeen time.Time

	ConnectedAt time.Time
}

func NewClient(

	conn *gws.Conn,

	userID uuid.UUID,

	deviceID uuid.UUID,

	hub *Hub,

) *Client {

	return &Client{

		ID: ClientID(uuid.New()),

		UserID: userID,

		DeviceID: deviceID,

		Conn: conn,

		Send: make(chan []byte, 256),

		Hub: hub,

		LastSeen: time.Now(),

		ConnectedAt: time.Now(),
	}
}