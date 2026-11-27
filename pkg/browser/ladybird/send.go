package ladybird

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coder/websocket/wsjson"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/message"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"time"
)

func (c *Client) send(
	method string,
	parameters any,
) (json.RawMessage, error) {
	c.mutex.Lock()
	connection := c.connection
	c.mutex.Unlock()

	if connection == nil {
		return nil, unavailable.Format(constant.LadybirdNotConnected)
	}

	identifier := int(c.identifier.Add(1))
	channel := make(chan *message.Reply, 1)
	c.mutex.Lock()
	c.pending[identifier] = channel
	c.mutex.Unlock()

	defer c.forget(identifier)
	e := wsjson.Write(
		context.Background(),
		connection,
		message.NewRequest(identifier, method, parameters),
	)

	if e != nil {
		return nil, fmt.Errorf("%s: %w", method, e)
	}

	select {
	case r := <-channel:
		return reply(method, r)
	case <-time.After(constant.LadybirdSendTimeout):
		return nil, timeout.Format("%s: no reply", method)
	}
}
