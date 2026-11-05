package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_group"

func (c *Client) TunnelGroups() ([]*tunnel_group.Group, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
