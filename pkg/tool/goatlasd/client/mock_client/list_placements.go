package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
)

func (c *Client) ListPlacements(
	_ context.Context,
	place string,
	source string,
	packageName string,
) ([]client.Placement, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []client.Placement

	for _, v := range c.PlacementList {
		if place != "" && (v.PlaceName == nil || *v.PlaceName != place) {
			continue
		}

		if source != "" && v.Source != source {
			continue
		}

		if packageName != "" && (v.Package == nil || *v.Package != packageName) {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
