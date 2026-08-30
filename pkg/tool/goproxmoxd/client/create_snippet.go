package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateSnippet(
	body client.CreateSnippetJSONRequestBody,
) (string, int) {
	result, e := c.client.CreateSnippetWithResponse(
		c.context,
		&client.CreateSnippetParams{Instance: &c.instance},
		body,
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
