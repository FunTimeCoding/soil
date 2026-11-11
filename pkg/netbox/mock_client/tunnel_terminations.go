package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"

func (c *Client) TunnelTerminations() ([]*tunnel_termination.Termination, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.tunnels, nil
}
