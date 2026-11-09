package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) TagDevices(_ string) ([]*device.Device, error) {
	return nil, c.Failure
}
