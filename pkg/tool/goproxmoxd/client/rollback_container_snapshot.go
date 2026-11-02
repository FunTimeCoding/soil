package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RollbackContainerSnapshot(
	x context.Context,
	identifier int64,
	name string,
	p *client.RollbackContainerSnapshotParams,
) (*client.TaskResult, error) {
	r, e := c.client.RollbackContainerSnapshotWithResponse(
		x,
		identifier,
		name,
		p,
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"rollback container snapshot",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
