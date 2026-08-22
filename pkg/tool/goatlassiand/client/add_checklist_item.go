package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddChecklistItem(
	key string,
	text string,
) string {
	result, e := c.client.AddChecklistItem(
		c.context,
		key,
		client.ChecklistTextRequest{Text: text},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
