package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateDevice(
	name string,
	role string,
	deviceType string,
	site string,
	tenant *string,
) *response.Response {
	result, e := c.client.CreateDevice(
		c.context,
		client.CreateDeviceRequest{
			Name:   name,
			Role:   role,
			Type:   deviceType,
			Site:   site,
			Tenant: tenant,
		},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
