package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetStatus(x context.Context) (*client.StatusResponse, error) {
	r, e := c.generated.GetStatusWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get status", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
