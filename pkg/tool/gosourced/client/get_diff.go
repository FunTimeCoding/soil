package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetDiff(
	x context.Context,
	module string,
	from string,
	to string,
) (client.DiffResponse, error) {
	p := &client.GetDiffParams{Module: module}

	if from != "" {
		p.From = &from
	}

	if to != "" {
		p.To = &to
	}

	r, e := c.client.GetDiffWithResponse(x, p)

	if e != nil {
		return client.DiffResponse{}, e
	}

	if r.JSON200 == nil {
		return client.DiffResponse{}, refusal.New(
			"diff",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
