package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) FeatureDefinitions(
	x context.Context,
) ([]*gitlab.FeatureDefinition, error) {
	result, _, e := c.client.Features.ListFeatureDefinitions(
		gitlab.WithContext(x),
	)

	return result, wrapError(e)
}
