package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Forwards(
	x context.Context,
	phrase string,
) ([]client.Forward, error) {
	r, e := c.client.ListForwardsWithResponse(
		x,
		&client.ListForwardsParams{Query: optional(phrase)},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("forwards", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
