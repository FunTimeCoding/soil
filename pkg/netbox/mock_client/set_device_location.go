package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceLocation(
	name string,
	locationName string,
) (*device.Device, error) {
	return c.changeDevice("SetDeviceLocation", name, locationName)
}
