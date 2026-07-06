package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/console_server_port"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (c *Client) ConsoleServerPorts() ([]*console_server_port.Port, error) {
	result, _, e := c.client.DcimAPI.DcimConsoleServerPortsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return console_server_port.NewSlice(result.Results), nil
}
