package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) Devices() ([]*device.Device, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return c.devices, nil
}
