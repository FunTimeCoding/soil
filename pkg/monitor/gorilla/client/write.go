package client

import (
	"github.com/gorilla/websocket"
	"log"
)

func (c *Client) Write(s string) {
	log.Printf("send: %s\n", s)

	if e := c.connection.WriteMessage(
		websocket.TextMessage,
		[]byte(s),
	); e != nil {
		log.Printf("write: %s\n", e)
	}
}
