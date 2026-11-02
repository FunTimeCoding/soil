package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ResetMachine(
	x context.Context,
	identifier int64,
	p *client.ResetMachineParams,
) (*client.TaskResult, error) {
	r, e := c.client.ResetMachineWithResponse(x, identifier, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("reset machine", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
