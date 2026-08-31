package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SearchPages(query string) *response.Response {
	result, e := c.client.SearchPages(
		c.context,
		&client.SearchPagesParams{Query: query},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
