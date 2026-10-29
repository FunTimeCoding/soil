package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) UnsetSerial(_ string) (*device.Device, error) {
	return nil, c.Failure
}
