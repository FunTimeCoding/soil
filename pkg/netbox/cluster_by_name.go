package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
)

func (c *Client) ClusterByName(n string) (*cluster.Cluster, error) {
	result, e := c.ClustersByName(n)

	if e != nil {
		return nil, e
	}

	for _, cl := range result {
		if cl.Name == n {
			return cl, nil
		}
	}

	return nil, not_found.New("cluster", n)
}
