package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/rack"
)

func (c *Client) MustRackByName(n string) *rack.Rack {
	result, e := c.RackByName(n)
	errors.PanicOnError(e)

	return result
}
