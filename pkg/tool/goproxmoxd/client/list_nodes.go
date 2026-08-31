package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListNodes() *response.Response {
	result, e := c.client.ListNodesWithResponse(
		c.context,
		&client.ListNodesParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
