package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) TransitionIssue(
	key string,
	transitionIdentifier string,
) string {
	result, e := c.client.TransitionIssue(
		c.context,
		key,
		client.TransitionRequest{TransitionIdentifier: transitionIdentifier},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
