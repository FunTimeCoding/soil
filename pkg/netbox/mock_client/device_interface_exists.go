package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) DeviceInterfaceExists(
	_ *device.Device,
	_ string,
) (bool, error) {
	return false, c.Failure
}
