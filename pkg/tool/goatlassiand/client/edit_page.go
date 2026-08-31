package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) EditPage(
	identifier string,
	oldText string,
	newText string,
	title string,
	message string,
	draft bool,
) *response.Response {
	body := client.EditPageJSONRequestBody{OldText: oldText, NewText: newText}

	if title != "" {
		body.Title = &title
	}

	if message != "" {
		body.Message = &message
	}

	if draft {
		body.Draft = &draft
	}

	result, e := c.client.EditPage(c.context, identifier, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
