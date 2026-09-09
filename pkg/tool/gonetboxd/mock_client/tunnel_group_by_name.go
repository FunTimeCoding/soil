package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_group"

func (c *Client) TunnelGroupByName(_ string) (*tunnel_group.Group, error) {
	return nil, nil
}
