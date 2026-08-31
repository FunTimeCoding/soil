package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) StopContainer(
	identifier int64,
	node *string,
) *response.Response {
	result, e := c.client.StopContainerWithResponse(
		c.context,
		identifier,
		&client.StopContainerParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
