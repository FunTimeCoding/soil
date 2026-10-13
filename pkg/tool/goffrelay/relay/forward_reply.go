package relay

import (
	"context"
	"github.com/coder/websocket"
	"log"
	"os"
)

func forwardReply(
	x context.Context,
	c *websocket.Conn,
	closed chan struct{},
) {
	defer func() { closed <- struct{}{} }()

	for {
		b, e := readFrame(os.Stdin)

		if e != nil {
			return
		}

		if f := c.Write(x, websocket.MessageText, b); f != nil {
			log.Printf("firefox relay: write: %v", f)

			return
		}
	}
}
