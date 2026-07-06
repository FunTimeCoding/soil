package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/console_server_port"
)

func (c *Client) MustConsoleServerPorts() []*console_server_port.Port {
	result, e := c.ConsoleServerPorts()
	errors.PanicOnError(e)

	return result
}
