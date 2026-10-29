package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListPlatforms(x context.Context) ([]*client.Platform, error) {
	r, e := c.client.ListPlatformsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list platforms", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
