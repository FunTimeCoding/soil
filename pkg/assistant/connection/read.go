package connection

import (
	"github.com/funtimecoding/soil/pkg/assistant/message"
	"github.com/gorilla/websocket"
)

func (c *Connection) Read() {
	for {
		select {
		case <-c.context.Done():
			return
		default:
			var m message.Message

			if e := c.connection.ReadJSON(&m); e != nil {
				if websocket.IsCloseError(
					e,
					websocket.CloseNormalClosure,
					websocket.CloseGoingAway,
				) {
					return
				}

				return
			}

			switch m.Type {
			case Result:
				c.consumeResult(&m)
			case Event:
				c.consumeEvent(&m)
			case Pong:
				// pong received
			}
		}
	}
}
