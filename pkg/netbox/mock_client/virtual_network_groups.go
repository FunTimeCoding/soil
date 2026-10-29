package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_network_group"

func (c *Client) VirtualNetworkGroups() ([]*virtual_network_group.Group, error) {
	return nil, c.Failure
}
