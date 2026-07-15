package websocket

import (
	"github.com/gorilla/websocket"
)

func (c *Client) WritePump() {

	defer c.Conn.Close()

	for {

		select {

		case message := <-c.Send:

			c.Conn.WriteMessage(
				websocket.TextMessage,
				message,
			)

		}
	}
}