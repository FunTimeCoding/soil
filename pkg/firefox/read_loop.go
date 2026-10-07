package firefox

import (
	"context"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
)

func (c *Client) readLoop(connection *websocket.Conn) {
	for {
		var r *message.Reply

		if e := wsjson.Read(context.Background(), connection, &r); e != nil {
			return
		}

		c.mutex.Lock()
		channel, okay := c.pending[r.Identifier]
		c.mutex.Unlock()

		if okay {
			channel <- r
		}
	}
}
