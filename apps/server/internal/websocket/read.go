package websocket

func (c *Client) ReadPump() {

	defer func() {

		c.Hub.Unregister <- c

		c.Conn.Close()

	}()

	for {

		var msg Message

		err := c.Conn.ReadJSON(&msg)

		if err != nil {

			break
		}
	}
}