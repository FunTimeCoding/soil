package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/console_port"

func (c *Client) ConsolePorts() ([]*console_port.Port, error) {
	return nil, c.Failure
}
