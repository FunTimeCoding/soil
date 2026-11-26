package ladybird

import (
	"context"
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Open() {
	if c.address == "" {
		c.launch()
	}

	c.waitDriver()
	s := c.createSession()
	connection, _, e := websocket.Dial(
		context.Background(),
		s.Value.Capabilities.WebSocket,
		&websocket.DialOptions{HTTPClient: c.pool.Client()},
	)
	errors.PanicOnError(e)
	connection.SetReadLimit(constant.ReadLimit)
	c.mutex.Lock()
	c.session = s.Value.SessionIdentifier
	c.connection = connection
	c.mutex.Unlock()
	go c.readLoop(connection)
}
