package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) RemoveDeviceLabel(
	device string,
	key string,
) (string, int) {
	result, e := c.client.RemoveDeviceLabel(c.context, device, key)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
