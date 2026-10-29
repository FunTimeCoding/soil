package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetPostReactions(
	x context.Context,
	identifier string,
) ([]string, error) {
	r, e := c.client.GetPostReactionsWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post reactions", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
