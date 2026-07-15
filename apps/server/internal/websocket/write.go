package websocket

import (
	"log"
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

			err := c.Conn.WriteMessage(
				gws.TextMessage,
				message,
			)

			if err != nil {

				log.Println(err)

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