package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListMachines(node *string) *response.Response {
	result, e := c.client.ListMachinesWithResponse(
		c.context,
		&client.ListMachinesParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
