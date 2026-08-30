package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ShutdownMachine(
	identifier int64,
	node *string,
) (string, int) {
	result, e := c.client.ShutdownMachineWithResponse(
		c.context,
		identifier,
		&client.ShutdownMachineParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
