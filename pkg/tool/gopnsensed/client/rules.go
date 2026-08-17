package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Rules(query *string) string {
	result, e := c.client.ListRules(
		c.context,
		&client.ListRulesParams{Query: query},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
