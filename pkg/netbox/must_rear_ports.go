package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/rear_port"
)

func (c *Client) MustRearPorts() []*rear_port.Port {
	result, e := c.RearPorts()
	errors.PanicOnError(e)

	return result
}
