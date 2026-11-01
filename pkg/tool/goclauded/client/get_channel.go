package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetChannel(
	x context.Context,
	params *client.GetChannelParams,
) (*client.CheckResponse, error) {
	r, e := c.generated.GetChannelWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get channel", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
