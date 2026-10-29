package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
)

func (c *Client) ListSightings(
	_ context.Context,
	place string,
	unclaimed bool,
) ([]client.Sighting, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []client.Sighting

	for _, v := range c.SightingList {
		claimed := v.PlaceKind != nil && *v.PlaceKind != ""

		if unclaimed && claimed {
			continue
		}

		if place != "" && (v.PlaceName == nil || *v.PlaceName != place) {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
