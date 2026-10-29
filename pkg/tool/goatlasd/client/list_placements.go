package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListPlacements(
	x context.Context,
	place string,
	source string,
	packageName string,
) ([]client.Placement, error) {
	r, e := c.client.ListPlacementsWithResponse(
		x,
		&client.ListPlacementsParams{
			Place:   optional(place),
			Source:  optional(source),
			Package: optional(packageName),
		},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list placements", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
