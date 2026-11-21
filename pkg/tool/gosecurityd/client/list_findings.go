package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListFindings(
	x context.Context,
	p *client.ListFindingsParams,
) ([]*client.FindingGroup, error) {
	r, e := c.client.ListFindingsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("findings", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
