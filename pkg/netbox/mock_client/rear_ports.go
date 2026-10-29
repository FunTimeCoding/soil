package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/rear_port"

func (c *Client) RearPorts() ([]*rear_port.Port, error) {
	return nil, c.Failure
}
