package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_group"

func (c *Client) CreateTunnelGroup(_ string) (*tunnel_group.Group, error) {
	return nil, nil
}
