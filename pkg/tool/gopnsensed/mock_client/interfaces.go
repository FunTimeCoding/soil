package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/network_interface"

func (c *Client) Interfaces() ([]*network_interface.Interface, error) {
	return nil, nil
}
