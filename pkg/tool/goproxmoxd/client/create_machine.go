package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateMachine(
	x context.Context,
	p *client.CreateMachineParams,
	b client.CreateMachineJSONRequestBody,
) (*client.CreateMachineResult, error) {
	r, e := c.client.CreateMachineWithResponse(x, p, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("create machine", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
