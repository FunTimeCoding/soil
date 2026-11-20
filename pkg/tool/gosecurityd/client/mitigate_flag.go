package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) MitigateFlag(
	x context.Context,
	identifier string,
	b client.MitigationRequest,
) (*client.Flag, error) {
	r, e := c.client.MitigateFlagWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("mitigate", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
