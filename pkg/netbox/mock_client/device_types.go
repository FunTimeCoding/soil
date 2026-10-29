package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device_type"

func (c *Client) DeviceTypes() ([]*device_type.Type, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return c.types, nil
}
