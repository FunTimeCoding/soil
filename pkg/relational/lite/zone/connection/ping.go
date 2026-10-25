package connection

import "context"

func (c *Connection) Ping(x context.Context) error {
	return c.inner.Ping(x)
}
