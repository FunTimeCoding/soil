package ladybird

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) stopDriver() {
	if c.driver == nil {
		return
	}

	errors.LogOnError(c.driver.Kill())
	c.driver = nil
}
