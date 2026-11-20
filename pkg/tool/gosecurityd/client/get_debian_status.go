package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetDebianStatus(
	x context.Context,
	identifier string,
	p *client.GetDebianStatusParams,
) (*client.DebianStatus, error) {
	r, e := c.client.GetDebianStatusWithResponse(x, identifier, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("debian status", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
