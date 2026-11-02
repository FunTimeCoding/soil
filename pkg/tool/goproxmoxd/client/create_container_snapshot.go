package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateContainerSnapshot(
	x context.Context,
	identifier int64,
	p *client.CreateContainerSnapshotParams,
	b client.CreateContainerSnapshotJSONRequestBody,
) (*client.TaskResult, error) {
	r, e := c.client.CreateContainerSnapshotWithResponse(x, identifier, p, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"create container snapshot",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
