package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) Reconcile(
	_ context.Context,
	p *client.ReconcileParams,
) (*client.ReconcilePlan, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result := *c.Plan
	result.Written = p != nil && p.Write != nil && *p.Write

	return &result, nil
}
