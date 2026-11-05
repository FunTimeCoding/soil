package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDevicePlatform(
	name string,
	platformName string,
) (*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.changeDevice("SetDevicePlatform", name, platformName)
}
