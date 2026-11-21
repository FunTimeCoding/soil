package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetScoreStatus(x context.Context) (*client.ScoreStatus, error) {
	r, e := c.client.GetScoreStatusWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("score status", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
