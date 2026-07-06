package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
)

func (c *Client) MustPowerOutlets() []*power_outlet.Outlet {
	result, e := c.PowerOutlets()
	errors.PanicOnError(e)

	return result
}
