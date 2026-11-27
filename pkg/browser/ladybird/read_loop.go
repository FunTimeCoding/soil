package ladybird

import (
	"context"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
)

func (c *Client) readLoop(connection *websocket.Conn) {
	for {
		r := message.NewReply()

		if e := wsjson.Read(context.Background(), connection, r); e != nil {
			return
		}

		if r.Type == constant.LadybirdEventType {
			c.event(r)

			continue
		}

		c.mutex.Lock()
		channel, okay := c.pending[r.Identifier]
		c.mutex.Unlock()

		if okay {
			channel <- r
		}
	}
}
