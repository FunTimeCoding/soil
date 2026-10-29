package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListTunnelTerminations(
	x context.Context,
) ([]*client.TunnelTermination, error) {
	r, e := c.client.ListTunnelTerminationsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"list tunnel terminations",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
