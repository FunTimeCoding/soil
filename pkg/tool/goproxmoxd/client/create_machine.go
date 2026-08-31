package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateMachine(
	body client.CreateMachineJSONRequestBody,
) *response.Response {
	result, e := c.client.CreateMachineWithResponse(
		c.context,
		&client.CreateMachineParams{Instance: &c.instance},
		body,
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
