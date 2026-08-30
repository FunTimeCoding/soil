package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) GetSnippet(name string) (string, int) {
	result, e := c.client.GetSnippetWithResponse(
		c.context,
		name,
		&client.GetSnippetParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
