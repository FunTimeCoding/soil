package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostCheck(
	x context.Context,
	b client.CheckRequest,
) (client.CheckResponse, error) {
	r, e := c.client.PostCheckWithResponse(x, b)

	if e != nil {
		return client.CheckResponse{}, e
	}

	if r.JSON200 == nil {
		return client.CheckResponse{}, refusal.New(
			"check",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
