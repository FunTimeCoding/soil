package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDevicePrimaryAddress(
	name string,
	address string,
) (*device.Device, error) {
	return c.changeDevice("SetDevicePrimaryAddress", name, address)
}
