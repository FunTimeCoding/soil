package connection

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/gorilla/websocket"
	"time"
)

func (c *Connection) Stop() {
	c.cancel()
	c.Lock()
	defer c.Unlock()

	if c.connection == nil {
		return
	}

	errors.PanicOnError(
		c.connection.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(
				websocket.CloseNormalClosure,
				"",
			),
		),
	)
	time.Sleep(time.Second)
	errors.LogClose(c.connection)
}
