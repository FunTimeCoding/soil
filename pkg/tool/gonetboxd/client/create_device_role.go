package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateDeviceRole(
	x context.Context,
	name string,
	virtualMachine bool,
) (client.DeviceRole, error) {
	r, e := c.client.CreateDeviceRoleWithResponse(
		x,
		client.CreateDeviceRoleRequest{
			Name:           name,
			VirtualMachine: &virtualMachine,
		},
	)

	if e != nil {
		return client.DeviceRole{}, e
	}

	if r.JSON201 == nil {
		return client.DeviceRole{}, refusal.New(
			"create device role",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
