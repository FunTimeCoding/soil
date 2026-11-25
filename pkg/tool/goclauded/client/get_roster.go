package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetRoster(
	x context.Context,
	params *client.GetRosterParams,
) (*client.RosterResponse, error) {
	r, e := c.generated.GetRosterWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get roster", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
