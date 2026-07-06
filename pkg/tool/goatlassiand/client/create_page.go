package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreatePage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	body string,
) string {
	result, e := c.client.CreatePage(
		c.context,
		client.CreatePageJSONRequestBody{
			SpaceIdentifier:  spaceIdentifier,
			ParentIdentifier: parentIdentifier,
			Title:            title,
			Body:             body,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
