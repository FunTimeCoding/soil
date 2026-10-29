package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) EditPage(
	x context.Context,
	identifier string,
	oldText string,
	newText string,
	title string,
	message string,
	draft bool,
) (*client.ConfluencePageDetail, error) {
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

	r, e := c.client.EditPageWithResponse(x, identifier, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("edit page", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
