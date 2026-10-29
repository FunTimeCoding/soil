package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateSite(
	x context.Context,
	name string,
) (client.Site, error) {
	r, e := c.client.CreateSiteWithResponse(
		x,
		client.CreateNameRequest{Name: name},
	)

	if e != nil {
		return client.Site{}, e
	}

	if r.JSON201 == nil {
		return client.Site{}, refusal.New("create site", r.HTTPResponse, r.Body)
	}

	return *r.JSON201, nil
}
