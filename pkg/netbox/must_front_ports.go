package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/front_port"
)

func (c *Client) MustFrontPorts() []*front_port.Port {
	result, e := c.FrontPorts()
	errors.PanicOnError(e)

	return result
}
