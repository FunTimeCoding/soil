package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListVirtualDisks(machine string) string {
	result, e := c.client.ListVirtualDisks(c.context, machine)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
