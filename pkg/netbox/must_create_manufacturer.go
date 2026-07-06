package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
)

func (c *Client) MustCreateManufacturer(name string) *manufacturer.Manufacturer {
	result, e := c.CreateManufacturer(name)
	errors.PanicOnError(e)

	return result
}
