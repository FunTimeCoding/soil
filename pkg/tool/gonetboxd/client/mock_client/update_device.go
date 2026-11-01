package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) UpdateDevice(
	_ context.Context,
	name string,
	body client.UpdateDeviceRequest,
) (client.Device, error) {
	if c.Failure != nil {
		return client.Device{}, c.Failure
	}

	if c.Updated == nil {
		c.Updated = map[string]client.UpdateDeviceRequest{}
	}

	c.Updated[name] = body

	return client.Device{}, nil
}
