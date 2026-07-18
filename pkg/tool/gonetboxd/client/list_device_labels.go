package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListDeviceLabels(device string) string {
	result, e := c.client.ListDeviceLabels(c.context, device)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
