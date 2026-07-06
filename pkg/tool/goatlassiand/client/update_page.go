package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdatePage(
	identifier string,
	title string,
	body string,
	message *string,
) string {
	result, e := c.client.UpdatePage(
		c.context,
		identifier,
		client.UpdatePageJSONRequestBody{
			Title:   title,
			Body:    body,
			Message: message,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
