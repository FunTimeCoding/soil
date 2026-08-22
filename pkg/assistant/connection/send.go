package connection

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
)

func (c *Connection) send(
	o command,
	s Subscriber,
) (uint64, error) {
	c.Lock()
	defer c.Unlock()

	if c.connection == nil {
		return 0, unreachable.Format("connection not open")
	}

	c.lastIdentifier++
	o.SetIdentifier(c.lastIdentifier)

	if s != nil {
		c.subscribers[c.lastIdentifier] = s
	}

	if e := c.connection.WriteJSON(o); e != nil {
		delete(c.subscribers, c.lastIdentifier)

		return 0, fmt.Errorf("send: %w", e)
	}

	return c.lastIdentifier, nil
}
