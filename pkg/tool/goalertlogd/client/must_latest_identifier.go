package client

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustLatestIdentifier() uint {
	result, e := c.LatestIdentifier()
	errors.PanicOnError(e)

	return result
}
