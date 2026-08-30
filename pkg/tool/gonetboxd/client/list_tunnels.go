package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListTunnels() (string, int) {
	result, e := c.client.ListTunnels(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
