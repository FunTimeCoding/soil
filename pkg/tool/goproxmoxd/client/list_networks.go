package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListNetworks(name string) (string, int) {
	result, e := c.client.ListNetworksWithResponse(
		c.context,
		name,
		&client.ListNetworksParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
