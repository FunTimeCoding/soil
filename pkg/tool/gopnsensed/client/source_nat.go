package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SourceNat(query *string) (string, int) {
	result, e := c.client.ListSourceNat(
		c.context,
		&client.ListSourceNatParams{Query: query},
	)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
