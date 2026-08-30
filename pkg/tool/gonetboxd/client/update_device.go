package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdateDevice(
	name string,
	body client.UpdateDeviceRequest,
) (string, int) {
	result, e := c.client.UpdateDevice(c.context, name, body)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
