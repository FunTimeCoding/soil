package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetEpisodes(
	x context.Context,
	identifier string,
) ([]*client.Item, error) {
	r, e := c.client.GetEpisodesWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get episodes", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
