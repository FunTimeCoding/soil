package connection

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/gorilla/websocket"
)

func (c *Connection) Close() {
	c.Lock()
	defer c.Unlock()

	if c.connection == nil {
		return
	}

	errors.LogOnError(
		c.connection.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		),
	)
	errors.LogClose(c.connection)
	c.connection = nil
}
