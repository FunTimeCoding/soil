package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessions(
	x context.Context,
	params *client.GetSessionsParams,
) (*client.SessionListResponse, error) {
	r, e := c.generated.GetSessionsWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get sessions", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
