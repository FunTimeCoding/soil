package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) StartContainer(
	identifier int64,
	node *string,
) *response.Response {
	result, e := c.client.StartContainerWithResponse(
		c.context,
		identifier,
		&client.StartContainerParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
