package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetPackages(
	x context.Context,
	p *client.GetPackagesParams,
) ([]client.Listing, error) {
	r, e := c.client.GetPackagesWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("packages", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
