package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SetDeviceTenant(
	_ string,
	_ string,
) (*device.Device, error) {
	return nil, nil
}
