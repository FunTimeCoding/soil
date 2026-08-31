package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddPageComment(
	identifier string,
	body string,
) *response.Response {
	result, e := c.client.AddPageComment(
		c.context,
		identifier,
		client.CommentRequest{Body: body},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
