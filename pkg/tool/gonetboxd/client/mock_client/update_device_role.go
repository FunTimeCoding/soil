package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) UpdateDeviceRole(
	_ context.Context,
	name string,
	body client.UpdateDeviceRoleRequest,
) (client.DeviceRole, error) {
	if c.Failure != nil {
		return client.DeviceRole{}, c.Failure
	}

	if c.UpdatedRoles == nil {
		c.UpdatedRoles = map[string]client.UpdateDeviceRoleRequest{}
	}

	c.UpdatedRoles[name] = body

	return client.DeviceRole{}, nil
}
