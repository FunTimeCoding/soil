package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) RemoveDeviceTag(
	device string,
	tag string,
) *response.Response {
	result, e := c.client.RemoveDeviceTag(c.context, device, tag)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
