package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListMachines(node *string) (string, int) {
	result, e := c.client.ListMachinesWithResponse(
		c.context,
		&client.ListMachinesParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
