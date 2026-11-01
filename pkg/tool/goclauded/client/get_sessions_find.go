package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionsFind(
	x context.Context,
	params *client.GetSessionsFindParams,
) (*client.FindResponse, error) {
	r, e := c.generated.GetSessionsFindWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get sessions find", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
