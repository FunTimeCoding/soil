package claim

import "context"

func (c *Connection) Open(x context.Context) error {
	b, e := c.client.Listen(x)

	if e != nil {
		return e
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.body = b

	return nil
}
