package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) GetNodeStatus(name string) (string, int) {
	result, e := c.client.GetNodeStatusWithResponse(
		c.context,
		name,
		&client.GetNodeStatusParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
