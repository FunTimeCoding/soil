package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
)

func (c *Client) MustManufacturers() []*manufacturer.Manufacturer {
	result, e := c.Manufacturers()
	errors.PanicOnError(e)

	return result
}
