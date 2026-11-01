package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSessionsSearch(
	x context.Context,
	params *client.GetSessionsSearchParams,
) (*client.SearchResponse, error) {
	r, e := c.generated.GetSessionsSearchWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get sessions search", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
