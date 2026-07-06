package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListTunnelGroups() string {
	result, e := c.client.ListTunnelGroups(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
