package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListPlaces(x context.Context) ([]client.Place, error) {
	r, e := c.client.ListPlacesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list places", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
