package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetResolve(
	x context.Context,
	params *client.GetResolveParams,
) (*client.ResolveResponse, []client.ResolveMatch, error) {
	r, e := c.generated.GetResolveWithResponse(x, params)

	if e != nil {
		return nil, nil, e
	}

	if r.JSON409 != nil {
		return nil, r.JSON409.Matches, nil
	}

	if r.JSON200 == nil {
		return nil, nil, refusal.New("get resolve", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil, nil
}
