package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/cluster_group"

func (c *Client) ClusterGroups() ([]*cluster_group.Group, error) {
	return nil, c.Failure
}
