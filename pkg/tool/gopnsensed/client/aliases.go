package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Aliases(query *string) string {
	result, e := c.client.ListAliases(
		c.context,
		&client.ListAliasesParams{Query: query},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
