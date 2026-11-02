package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateMachine(
	x context.Context,
	identifier int64,
	p *client.UpdateMachineParams,
	b client.UpdateMachineJSONRequestBody,
) (*client.StatusResult, error) {
	r, e := c.client.UpdateMachineWithResponse(x, identifier, p, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("update machine", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
