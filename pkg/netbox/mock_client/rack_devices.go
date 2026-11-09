package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) RackDevices(_ int32) ([]*device.Device, error) {
	return nil, c.Failure
}
