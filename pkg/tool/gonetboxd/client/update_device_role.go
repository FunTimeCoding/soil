package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateDeviceRole(
	x context.Context,
	name string,
	body client.UpdateDeviceRoleRequest,
) (client.DeviceRole, error) {
	r, e := c.client.UpdateDeviceRoleWithResponse(x, name, body)

	if e != nil {
		return client.DeviceRole{}, e
	}

	if r.JSON200 == nil {
		return client.DeviceRole{}, refusal.New(
			"update device role",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
