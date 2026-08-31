package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Blocklists(query *string) *response.Response {
	result, e := c.client.ListBlocklists(
		c.context,
		&client.ListBlocklistsParams{Query: query},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
