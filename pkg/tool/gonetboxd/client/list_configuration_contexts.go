package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListConfigurationContexts(
	x context.Context,
) ([]*client.ConfigurationContext, error) {
	r, e := c.client.ListConfigurationContextsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"list configuration contexts",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
