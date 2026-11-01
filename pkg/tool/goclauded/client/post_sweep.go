package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostSweep(x context.Context) (*client.SweepResponse, error) {
	r, e := c.generated.PostSweepWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post sweep", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
