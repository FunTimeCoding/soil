package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_chassis"

func (c *Client) VirtualChassis() ([]*virtual_chassis.Chassis, error) {
	return nil, c.Failure
}
