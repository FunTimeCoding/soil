package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/power_outlet"

func (c *Client) PowerOutlets() ([]*power_outlet.Outlet, error) {
	return nil, c.Failure
}
