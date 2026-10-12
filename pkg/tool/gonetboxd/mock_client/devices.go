package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) Devices() ([]*device.Device, error) {
	return c.devices, nil
}
