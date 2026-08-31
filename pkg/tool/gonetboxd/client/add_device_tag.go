package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddDeviceTag(
	device string,
	tag string,
) *response.Response {
	result, e := c.client.AddDeviceTag(c.context, device, tag)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
