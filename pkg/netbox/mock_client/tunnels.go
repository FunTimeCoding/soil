package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel"

func (c *Client) Tunnels() ([]*tunnel.Tunnel, error) {
	return nil, nil
}
