package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceDescription(
	name string,
	description string,
) (*device.Device, error) {
	return c.changeDevice("SetDeviceDescription", name, description)
}
