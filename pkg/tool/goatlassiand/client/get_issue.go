package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetIssue(
	key string,
	comments bool,
) *response.Response {
	params := &client.GetIssueParams{}

	if comments {
		params.Comments = &comments
	}

	result, e := c.client.GetIssue(c.context, key, params)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
