package connection

import "context"

func (c *Connection) Open(x context.Context) error {
	e := c.connect(x)

	if e == nil {
		e = c.Subscribe("", c.subscriber)
	}

	if e != nil {
		c.Close()
	}

	return e
}
