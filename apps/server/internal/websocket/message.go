package websocket

import "encoding/json"

const (
	MessageHeartbeat = "heartbeat"
	MessagePong      = "pong"
)

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}