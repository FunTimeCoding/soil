package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/power_port"

func (c *Client) DevicePowerPorts(device int32) ([]*power_port.Port, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.ports[device], nil
}
