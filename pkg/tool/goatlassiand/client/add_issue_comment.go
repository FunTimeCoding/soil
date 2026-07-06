package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddIssueComment(
	key string,
	body string,
) string {
	result, e := c.client.AddIssueComment(
		c.context,
		key,
		client.CommentRequest{Body: body},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
