package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) RemoveDeviceTag(
	device string,
	tag string,
) string {
	result, e := c.client.RemoveDeviceTag(c.context, device, tag)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
