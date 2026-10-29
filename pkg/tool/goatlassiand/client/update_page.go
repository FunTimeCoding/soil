package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdatePage(
	x context.Context,
	identifier string,
	title string,
	body string,
	message *string,
) (*client.ConfluencePageDetail, error) {
	r, e := c.client.UpdatePageWithResponse(
		x,
		identifier,
		client.UpdatePageJSONRequestBody{
			Title:   title,
			Body:    body,
			Message: message,
		},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("update page", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
