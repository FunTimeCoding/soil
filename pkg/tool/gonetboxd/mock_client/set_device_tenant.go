package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceTenant(
	name string,
	tenantName string,
) (*device.Device, error) {
	return c.changeDevice("SetDeviceTenant", name, tenantName)
}
