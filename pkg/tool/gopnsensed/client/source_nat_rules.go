package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SourceNatRules(
	x context.Context,
	phrase string,
) ([]client.SourceNatRule, error) {
	r, e := c.client.ListSourceNatWithResponse(
		x,
		&client.ListSourceNatParams{Query: optional(phrase)},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("source nat rules", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
