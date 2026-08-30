package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListSnippets() (string, int) {
	result, e := c.client.ListSnippetsWithResponse(
		c.context,
		&client.ListSnippetsParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
