package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) GetContainer(
	identifier int64,
	node *string,
) (string, int) {
	result, e := c.client.GetContainerWithResponse(
		c.context,
		identifier,
		&client.GetContainerParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
