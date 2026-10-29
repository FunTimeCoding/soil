package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Rules(
	x context.Context,
	phrase string,
) ([]client.Rule, error) {
	r, e := c.client.ListRulesWithResponse(
		x,
		&client.ListRulesParams{Query: optional(phrase)},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("rules", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
