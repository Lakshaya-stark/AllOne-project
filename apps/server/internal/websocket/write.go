package websocket

import (
	
	"time"

	gws "github.com/gorilla/websocket"
)

func (c *Client) WritePump() {

	ticker := time.NewTicker(PingPeriod)

	defer func() {

		ticker.Stop()

		c.Conn.Close()

	}()

	for {

	select {

	case <-c.Context.Done():

		return

	case message, ok := <-c.Send:

		c.Conn.SetWriteDeadline(
			time.Now().Add(WriteWait),
		)

		if !ok {

			c.Conn.WriteMessage(
				gws.CloseMessage,
				[]byte{},
			)

			return
		}

		if err := c.Conn.WriteMessage(
			gws.TextMessage,
			message,
		); err != nil {

			return
		}

	case <-ticker.C:

		c.Conn.SetWriteDeadline(
			time.Now().Add(WriteWait),
		)

		if err := c.Conn.WriteMessage(
			gws.PingMessage,
			nil,
		); err != nil {

			return
		}
	}
}
}