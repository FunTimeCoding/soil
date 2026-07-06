package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
)

func (c *Client) MustRacks() []*rack.Rack {
	result, e := c.Racks()
	errors.PanicOnError(e)

	return result
}
