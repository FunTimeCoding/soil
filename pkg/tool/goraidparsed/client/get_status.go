package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetStatus(x context.Context) (client.StatusResponse, error) {
	r, e := c.client.GetStatusWithResponse(x)

	if e != nil {
		return client.StatusResponse{}, e
	}

	if r.JSON200 == nil {
		return client.StatusResponse{}, refusal.New(
			"get status",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
