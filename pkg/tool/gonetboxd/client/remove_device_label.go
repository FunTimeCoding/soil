package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) RemoveDeviceLabel(
	device string,
	key string,
) {
	result, e := c.client.RemoveDeviceLabel(c.context, device, key)
	errors.PanicOnError(e)
	web.ReadString(result)
}
