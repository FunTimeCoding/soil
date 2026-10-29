package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/front_port"

func (c *Client) FrontPorts() ([]*front_port.Port, error) {
	return nil, c.Failure
}
