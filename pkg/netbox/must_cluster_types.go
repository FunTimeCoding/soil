package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
)

func (c *Client) MustClusterTypes() []*cluster_type.Type {
	result, e := c.ClusterTypes()
	errors.PanicOnError(e)

	return result
}
