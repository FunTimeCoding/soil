package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
)

func (c *Client) MustCreateClusterType(name string) *cluster_type.Type {
	result, e := c.CreateClusterType(name)
	errors.PanicOnError(e)

	return result
}
