package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel"

func (c *Client) Tunnel(_ string) (*tunnel.Tunnel, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
