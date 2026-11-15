package assistant_server

import "github.com/gorilla/websocket"

func say(
	c *websocket.Conn,
	notation string,
) error {
	return c.WriteMessage(websocket.TextMessage, []byte(notation))
}
