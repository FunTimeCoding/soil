package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/location"
)

func (c *Client) MustLocations() []*location.Location {
	result, e := c.Locations()
	errors.PanicOnError(e)

	return result
}
