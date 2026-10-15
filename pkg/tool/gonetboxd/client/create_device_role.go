package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateDeviceRole(
	name string,
	virtualMachine bool,
) *response.Response {
	result, e := c.client.CreateDeviceRole(
		c.context,
		client.CreateDeviceRoleRequest{
			Name:           name,
			VirtualMachine: &virtualMachine,
		},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
