package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SearchPages(query string) string {
	result, e := c.client.SearchPages(
		c.context,
		&client.SearchPagesParams{Query: query},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
