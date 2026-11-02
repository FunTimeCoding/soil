package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListMachineSnapshots(
	x context.Context,
	identifier int64,
	p *client.ListMachineSnapshotsParams,
) ([]client.Snapshot, error) {
	r, e := c.client.ListMachineSnapshotsWithResponse(x, identifier, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"list machine snapshots",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
