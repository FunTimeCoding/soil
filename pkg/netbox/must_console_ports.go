package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/console_port"
)

func (c *Client) MustConsolePorts() []*console_port.Port {
	result, e := c.ConsolePorts()
	errors.PanicOnError(e)

	return result
}
