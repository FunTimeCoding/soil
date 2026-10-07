package firefox

import (
	"context"
	"fmt"
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/firefox/types/message"
	"time"
)

func (c *Client) send(
	method string,
	parameters any,
) (*message.Reply, error) {
	c.mutex.Lock()
	connection := c.connection
	c.mutex.Unlock()

	if connection == nil {
		return message.NewReply(), unreachable.Format("extension not connected")
	}

	identifier := int(c.identifier.Add(1))
	channel := make(chan *message.Reply, 1)
	c.mutex.Lock()
	c.pending[identifier] = channel
	c.mutex.Unlock()

	defer func() {
		c.mutex.Lock()
		delete(c.pending, identifier)
		c.mutex.Unlock()
	}()
	e := wsjson.Write(
		context.Background(),
		connection,
		message.NewRequest(method, parameters, identifier))

	if e != nil {
		return message.NewReply(), fmt.Errorf("%s: %w", method, e)
	}

	select {
	case r := <-channel:
		if r.Error != "" {
			return message.NewReply(), fmt.Errorf("%s: %s", method, r.Error)
		}

		return r, nil
	case <-time.After(10 * time.Second):
		return message.NewReply(), timeout.Format(
			"%s: timeout waiting for reply",
			method,
		)
	}
}
