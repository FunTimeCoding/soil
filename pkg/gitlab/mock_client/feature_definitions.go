package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) FeatureDefinitions(
	_ context.Context,
) ([]*gitlab.FeatureDefinition, error) {
	return nil, c.failure
}
