package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListNodes() (string, int) {
	result, e := c.client.ListNodesWithResponse(
		c.context,
		&client.ListNodesParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
