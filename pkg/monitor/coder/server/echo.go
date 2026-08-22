package server

import (
	"context"
	"fmt"
	"github.com/coder/websocket"
	"golang.org/x/time/rate"
	"io"
	"time"
)

// echo reads from the WebSocket connection and then writes
// the received message back to it.
// The entire function has 10s to complete.
func echo(
	c *websocket.Conn,
	l *rate.Limiter,
) error {
	x, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if e := l.Wait(x); e != nil {
		return e
	}

	messageType, r, readFail := c.Reader(x)

	if readFail != nil {
		return readFail
	}

	w, writeFail := c.Writer(x, messageType)

	if writeFail != nil {
		return writeFail
	}

	if _, e := io.Copy(w, r); e != nil {
		return fmt.Errorf("copy: %w", e)
	}

	return w.Close()
}
