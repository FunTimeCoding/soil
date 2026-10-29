package firefox

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) Close() {
	c.mutex.Lock()
	connection := c.connection
	c.connection = nil
	c.closed = true
	c.mutex.Unlock()

	if connection != nil {
		errors.LogOnError(connection.CloseNow())
	}
}
