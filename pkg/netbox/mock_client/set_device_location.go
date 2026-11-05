package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceLocation(
	name string,
	locationName string,
) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.changeDevice("SetDeviceLocation", name, locationName)
}
