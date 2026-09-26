package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
)

func (c *Client) Submit(intent string) int {
	result, e := c.client.PostRunWithResponse(
		c.context,
		client.RunRequest{Intent: intent},
	)
	errors.PanicOnError(e)

	return result.StatusCode()
}
