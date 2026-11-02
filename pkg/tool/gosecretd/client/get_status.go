package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetStatus(x context.Context) (*client.Status, error) {
	r, e := c.client.GetStatusWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get status", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
