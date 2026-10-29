package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Processes(x context.Context) ([]client.Process, error) {
	r, e := c.client.ListProcessesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("processes", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
