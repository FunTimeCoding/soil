package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel"

func (c *Client) TunnelByName(_ string) (*tunnel.Tunnel, error) {
	return nil, nil
}
