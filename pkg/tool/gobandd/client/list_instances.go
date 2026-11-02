package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListInstances(x context.Context) ([]client.Instance, error) {
	r, e := c.client.ListInstancesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("instances", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
