package client

import (
	"github.com/gorilla/websocket"
	"log"
)

func (c *Client) Send(s string) {
	if f := c.Connection.WriteMessage(
		websocket.TextMessage,
		[]byte(s),
	); f != nil {
		log.Printf("write: %s\n", f)
	}
}
