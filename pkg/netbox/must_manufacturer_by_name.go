package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
)

func (c *Client) MustManufacturerByName(n string) *manufacturer.Manufacturer {
	result, e := c.ManufacturerByName(n)
	errors.PanicOnError(e)

	return result
}
