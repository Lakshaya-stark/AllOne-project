package websocket

import (
	"encoding/json"
	"log"
	"time"

	gws "github.com/gorilla/websocket"
)

func (c *Client) ReadPump() {

	defer func() {

		c.Hub.Unregister(c)

		c.Conn.Close()

	}()

	c.Conn.SetReadLimit(MaxMessageSize)

	c.Conn.SetReadDeadline(
		time.Now().Add(PongWait),
	)

	c.Conn.SetPongHandler(func(string) error {

		c.LastSeen = time.Now()

		c.Conn.SetReadDeadline(
			time.Now().Add(PongWait),
		)

		return nil
	})

	for {

		var msg Message

		err := c.Conn.ReadJSON(&msg)

		if err != nil {

			if gws.IsUnexpectedCloseError(
				err,
				gws.CloseGoingAway,
				gws.CloseAbnormalClosure,
			) {

				log.Println(err)
			}

			break
		}

		switch msg.Type {

		case MessageHeartbeat:

			res, _ := json.Marshal(Message{
				Type: MessagePong,
			})

			c.Send <- res

		default:

			log.Println("Received:", msg.Type)
		}
	}
}