package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetIssue(
	key string,
	comments bool,
) (string, int) {
	params := &client.GetIssueParams{}

	if comments {
		params.Comments = &comments
	}

	result, e := c.client.GetIssue(c.context, key, params)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
