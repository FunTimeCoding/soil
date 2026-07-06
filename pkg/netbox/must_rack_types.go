package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/rack_type"
)

func (c *Client) MustRackTypes() []*rack_type.Type {
	result, e := c.RackTypes()
	errors.PanicOnError(e)

	return result
}
