package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddVirtualTag(
	name string,
	tag string,
) string {
	result, e := c.client.AddVirtualTag(c.context, name, tag)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
