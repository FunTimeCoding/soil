package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/wireless_network_group"

func (c *Client) WirelessNetworkGroups() ([]*wireless_network_group.Group, error) {
	return nil, c.Failure
}
