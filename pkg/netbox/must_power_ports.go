package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
)

func (c *Client) MustPowerPorts() []*power_port.Port {
	result, e := c.PowerPorts()
	errors.PanicOnError(e)

	return result
}
