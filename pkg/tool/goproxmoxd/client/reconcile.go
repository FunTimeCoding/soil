package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Reconcile(
	x context.Context,
	p *client.ReconcileParams,
) (*client.ReconcilePlan, error) {
	r, e := c.client.ReconcileWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(constant.Reconcile, r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
