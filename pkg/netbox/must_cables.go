package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cable"
)

func (c *Client) MustCables() []*cable.Cable {
	result, e := c.Cables()
	errors.PanicOnError(e)

	return result
}
