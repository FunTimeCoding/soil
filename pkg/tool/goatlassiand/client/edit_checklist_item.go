package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) EditChecklistItem(
	key string,
	index int,
	text string,
) *response.Response {
	result, e := c.client.EditChecklistItem(
		c.context,
		key,
		index,
		client.ChecklistTextRequest{Text: text},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
