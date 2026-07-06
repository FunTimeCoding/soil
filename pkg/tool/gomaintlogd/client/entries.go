package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Entries() string {
	result, e := c.client.GetEntries(c.context, &client.GetEntriesParams{})
	errors.PanicOnError(e)

	return web.ReadString(result)
}
