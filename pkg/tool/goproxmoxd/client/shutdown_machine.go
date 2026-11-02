package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ShutdownMachine(
	x context.Context,
	identifier int64,
	p *client.ShutdownMachineParams,
) (*client.TaskResult, error) {
	r, e := c.client.ShutdownMachineWithResponse(x, identifier, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("shutdown machine", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
