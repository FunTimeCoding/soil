package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Leases(query *string) (string, int) {
	result, e := c.client.ListLeases(
		c.context,
		&client.ListLeasesParams{Query: query},
	)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
