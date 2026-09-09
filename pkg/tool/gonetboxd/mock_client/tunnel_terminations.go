package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"

func (c *Client) TunnelTerminations() ([]*tunnel_termination.Termination, error) {
	return nil, nil
}
