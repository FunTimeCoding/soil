package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CloneMachine(
	identifier int64,
	node *string,
	body client.CloneMachineJSONRequestBody,
) (string, int) {
	result, e := c.client.CloneMachineWithResponse(
		c.context,
		identifier,
		&client.CloneMachineParams{Instance: &c.instance, Node: node},
		body,
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
