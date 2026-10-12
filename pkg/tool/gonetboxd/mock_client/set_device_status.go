package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceStatus(
	name string,
	status string,
) (*device.Device, error) {
	return c.changeDevice("SetDeviceStatus", name, status)
}
