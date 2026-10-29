package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListSightings(
	x context.Context,
	place string,
	unclaimed bool,
) ([]client.Sighting, error) {
	p := &client.ListSightingsParams{Place: optional(place)}

	if unclaimed {
		p.Unclaimed = &unclaimed
	}

	r, e := c.client.ListSightingsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list sightings", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
