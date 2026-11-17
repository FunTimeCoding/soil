package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/power_outlet"

func (c *Client) DevicePowerOutlets(
	device int32,
) ([]*power_outlet.Outlet, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.outlets[device], nil
}
