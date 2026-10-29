package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListSites(x context.Context) ([]*client.Site, error) {
	r, e := c.client.ListSitesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list sites", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
