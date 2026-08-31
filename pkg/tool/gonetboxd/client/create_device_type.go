package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateDeviceType(
	model string,
	manufacturer string,
) *response.Response {
	result, e := c.client.CreateDeviceType(
		c.context,
		client.CreateDeviceTypeRequest{
			Model:        model,
			Manufacturer: manufacturer,
		},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
