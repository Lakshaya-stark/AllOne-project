package websocket

import (
	"github.com/google/uuid"
	"encoding/json")

type Message struct {
	ID string `json:"id,omitempty"`

	Type MessageType `json:"type"`

	From *uuid.UUID `json:"from,omitempty"`

	To *uuid.UUID `json:"to,omitempty"`

	Payload json.RawMessage `json:"payload,omitempty"`
}