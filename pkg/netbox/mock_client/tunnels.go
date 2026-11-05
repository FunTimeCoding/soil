package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel"

func (c *Client) Tunnels() ([]*tunnel.Tunnel, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
