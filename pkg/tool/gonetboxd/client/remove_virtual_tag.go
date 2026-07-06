package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) RemoveVirtualTag(
	name string,
	tag string,
) string {
	result, e := c.client.RemoveVirtualTag(c.context, name, tag)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
