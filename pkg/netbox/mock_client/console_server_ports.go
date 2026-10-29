package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/console_server_port"

func (c *Client) ConsoleServerPorts() ([]*console_server_port.Port, error) {
	return nil, c.Failure
}
