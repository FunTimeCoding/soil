package claim

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Connection) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.body == nil {
		return
	}

	errors.LogClose(c.body)
	c.body = nil
}
