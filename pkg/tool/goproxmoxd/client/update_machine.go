package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) UpdateMachine(
	identifier int64,
	node *string,
	body client.UpdateMachineJSONRequestBody,
) (string, int) {
	result, e := c.client.UpdateMachineWithResponse(
		c.context,
		identifier,
		&client.UpdateMachineParams{Instance: &c.instance, Node: node},
		body,
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
