package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateMachine(
	body client.CreateMachineJSONRequestBody,
) string {
	result, e := c.client.CreateMachineWithResponse(
		c.context,
		&client.CreateMachineParams{
			Instance: &c.instance,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
