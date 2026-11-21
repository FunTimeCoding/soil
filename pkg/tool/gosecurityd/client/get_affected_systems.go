package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetAffectedSystems(
	x context.Context,
	identifier string,
) (*client.AffectedSystems, error) {
	r, e := c.client.GetAffectedSystemsWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("affected systems", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
