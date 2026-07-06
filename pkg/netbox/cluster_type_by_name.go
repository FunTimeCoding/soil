package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
)

func (c *Client) ClusterTypeByName(n string) (*cluster_type.Type, error) {
	result, e := c.ClusterTypes()

	if e != nil {
		return nil, e
	}

	for _, t := range result {
		if t.Name == n {
			return t, nil
		}
	}

	return nil, fmt.Errorf("cluster type not found: %s", n)
}
