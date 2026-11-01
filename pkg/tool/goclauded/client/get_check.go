package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetCheck(
	x context.Context,
	params *client.GetCheckParams,
) (*client.CheckResponse, error) {
	r, e := c.generated.GetCheckWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get check", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
