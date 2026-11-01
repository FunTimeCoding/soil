package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetProfile(
	x context.Context,
	p *client.GetProfileParams,
) (*client.ProfileResponse, error) {
	r, e := c.client.GetProfileWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get profile", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
