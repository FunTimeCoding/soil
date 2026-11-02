package claim

import "context"

func (c *Connection) deliver(m Message) {
	c.mutex.Lock()
	x := c.context
	c.mutex.Unlock()

	if x == nil {
		x = context.Background()
	}

	select {
	case c.updates <- m:
	case <-x.Done():
	}
}
