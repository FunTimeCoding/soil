package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetPageChildren(
	x context.Context,
	identifier string,
) ([]*client.ConfluencePage, error) {
	r, e := c.client.GetPageChildrenWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("page children", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
