package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/power_port"

func (c *Client) PowerPorts() ([]*power_port.Port, error) {
	return nil, c.Failure
}
