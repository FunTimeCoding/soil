package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListContainerSnapshots(
	x context.Context,
	identifier int64,
	p *client.ListContainerSnapshotsParams,
) ([]client.Snapshot, error) {
	r, e := c.client.ListContainerSnapshotsWithResponse(x, identifier, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"list container snapshots",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
