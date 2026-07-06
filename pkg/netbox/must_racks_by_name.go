package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
)

func (c *Client) MustRacksByName(n string) []*rack.Rack {
	result, e := c.RacksByName(n)
	errors.PanicOnError(e)

	return result
}
