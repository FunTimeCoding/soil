package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SearchUsers(query string) (string, int) {
	result, e := c.client.SearchUsers(
		c.context,
		&client.SearchUsersParams{Query: query},
	)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
