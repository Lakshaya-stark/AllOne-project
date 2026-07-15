package websocket

import "github.com/google/uuid"

type ClientID uuid.UUID

type DeviceID uuid.UUID

type UserID uuid.UUID

type MessageType string

const (

	MessageHeartbeat MessageType = "heartbeat"

	MessagePong MessageType = "pong"

	MessageClipboard MessageType = "clipboard"

	MessageFile MessageType = "file"

	MessageNotification MessageType = "notification"

	MessageBroadcast MessageType = "broadcast"
)