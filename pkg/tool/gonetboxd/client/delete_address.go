package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteAddress(identifier int32) (string, int) {
	result, e := c.client.DeleteAddress(c.context, identifier)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
