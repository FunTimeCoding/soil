package relay

import (
	"context"
	"github.com/coder/websocket"
	"log"
	"os"
)

func forwardRequest(
	x context.Context,
	c *websocket.Conn,
	closed chan struct{},
) {
	defer func() { closed <- struct{}{} }()

	for {
		_, b, e := c.Read(x)

		if e != nil {
			return
		}

		if f := writeFrame(os.Stdout, b); f != nil {
			log.Printf("firefox relay: frame: %v", f)

			return
		}
	}
}
