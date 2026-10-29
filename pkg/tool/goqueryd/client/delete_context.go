package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteContext(
	x context.Context,
	p *client.DeleteContextParams,
) (bool, error) {
	r, e := c.client.DeleteContextWithResponse(x, p)

	if e != nil {
		return false, e
	}

	if r.JSON200 == nil {
		return false, refusal.New("remove context", r.HTTPResponse, r.Body)
	}

	return r.JSON200.Deleted != nil && *r.JSON200.Deleted, nil
}
