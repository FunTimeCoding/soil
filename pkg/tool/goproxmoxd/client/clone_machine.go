package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CloneMachine(
	x context.Context,
	identifier int64,
	p *client.CloneMachineParams,
	b client.CloneMachineJSONRequestBody,
) (*client.CreateMachineResult, error) {
	r, e := c.client.CloneMachineWithResponse(x, identifier, p, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("clone machine", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
