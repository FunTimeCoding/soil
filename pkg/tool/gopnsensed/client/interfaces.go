package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Interfaces() string {
	result, e := c.client.ListInterfaces(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
