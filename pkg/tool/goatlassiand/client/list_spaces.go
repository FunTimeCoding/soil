package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListSpaces(x context.Context) ([]*client.ConfluenceSpace, error) {
	r, e := c.client.ListSpacesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("spaces", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
