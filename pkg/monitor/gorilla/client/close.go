package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func (c *Client) Close() {
	defer errors.LogClose(c.connection)

	if e := c.connection.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	); e != nil {
		log.Printf("write close: %s\n", e)

		return
	}

	select {
	case <-c.done:
		log.Printf("close done")
	case <-time.After(time.Second):
		log.Printf("close timeout")
	}
}
