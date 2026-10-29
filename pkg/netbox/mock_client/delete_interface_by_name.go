package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) DeleteInterfaceByName(
	_ *device.Device,
	_ string,
) error {
	return c.Failure
}
