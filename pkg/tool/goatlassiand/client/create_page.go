package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreatePage(
	x context.Context,
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	body string,
) (*client.ConfluencePageDetail, error) {
	r, e := c.client.CreatePageWithResponse(
		x,
		client.CreatePageJSONRequestBody{
			SpaceIdentifier:  spaceIdentifier,
			ParentIdentifier: parentIdentifier,
			Title:            title,
			Body:             body,
		},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON201 == nil {
		return nil, refusal.New("create page", r.HTTPResponse, r.Body)
	}

	return r.JSON201, nil
}
