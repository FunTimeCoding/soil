package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) UpdateMachine(
	identifier int64,
	node *string,
	body client.UpdateMachineJSONRequestBody,
) *response.Response {
	result, e := c.client.UpdateMachineWithResponse(
		c.context,
		identifier,
		&client.UpdateMachineParams{Instance: &c.instance, Node: node},
		body,
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
