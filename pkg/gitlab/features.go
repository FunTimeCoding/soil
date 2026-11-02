package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Features(x context.Context) ([]*gitlab.Feature, error) {
	result, _, e := c.client.Features.ListFeatures(gitlab.WithContext(x))

	return result, wrapError(e)
}
