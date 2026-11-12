package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostIndexPrune(x context.Context) ([]string, error) {
	r, e := c.client.PostIndexPruneWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("index prune", r.HTTPResponse, r.Body)
	}

	return r.JSON200.Removed, nil
}
