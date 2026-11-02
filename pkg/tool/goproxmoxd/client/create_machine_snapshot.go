package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateMachineSnapshot(
	x context.Context,
	identifier int64,
	p *client.CreateMachineSnapshotParams,
	b client.CreateMachineSnapshotJSONRequestBody,
) (*client.TaskResult, error) {
	r, e := c.client.CreateMachineSnapshotWithResponse(x, identifier, p, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"create machine snapshot",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
