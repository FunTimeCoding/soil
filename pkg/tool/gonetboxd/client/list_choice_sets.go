package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListChoiceSets(
	x context.Context,
) ([]*client.ChoiceSet, error) {
	r, e := c.client.ListChoiceSetsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list choice sets", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
