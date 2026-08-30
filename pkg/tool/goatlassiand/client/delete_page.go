package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeletePage(
	identifier string,
	draft bool,
) (string, int) {
	params := &client.DeletePageParams{}

	if draft {
		params.Draft = &draft
	}

	result, e := c.client.DeletePage(c.context, identifier, params)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
