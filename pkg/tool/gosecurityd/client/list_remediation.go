package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListRemediation(
	x context.Context,
	p *client.ListRemediationParams,
) (*client.RemediationList, error) {
	r, e := c.client.ListRemediationWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("remediation", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
