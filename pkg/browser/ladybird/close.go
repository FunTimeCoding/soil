package ladybird

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) Close() {
	c.mutex.Lock()
	connection := c.connection
	session := c.session
	c.session = ""
	c.connection = nil
	c.mutex.Unlock()
	defer c.pool.Close()
	defer c.stopDriver()

	if connection != nil {
		errors.LogOnError(connection.CloseNow())
	}

	if session != "" {
		errors.LogOnError(c.endSession(session))
	}
}
