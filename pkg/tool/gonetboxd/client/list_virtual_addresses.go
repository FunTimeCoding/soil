package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListVirtualAddresses(name string) string {
	result, e := c.client.ListVirtualAddresses(c.context, name)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
