package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SearchIssues(
	query string,
	limit *int,
) string {
	result, e := c.client.SearchIssues(
		c.context,
		&client.SearchIssuesParams{Query: query, Limit: limit},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
