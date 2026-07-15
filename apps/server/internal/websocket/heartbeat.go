package websocket

import (
	"encoding/json"
)

func (c *Client) HandleHeartbeat() error {

	msg := Message{
		Type: MessagePong,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	c.Send <- data

	return nil
}