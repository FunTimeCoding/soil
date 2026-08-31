package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DeleteMachine(
	identifier int64,
	node *string,
	purge *bool,
) *response.Response {
	result, e := c.client.DeleteMachineWithResponse(
		c.context,
		identifier,
		&client.DeleteMachineParams{
			Instance: &c.instance,
			Node:     node,
			Purge:    purge,
		},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
