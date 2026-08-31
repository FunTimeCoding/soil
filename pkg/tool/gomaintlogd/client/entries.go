package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Entries() *response.Response {
	result, e := c.client.GetEntries(c.context, &client.GetEntriesParams{})
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
