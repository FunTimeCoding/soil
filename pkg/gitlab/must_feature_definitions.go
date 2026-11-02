package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustFeatureDefinitions(
	x context.Context,
) []*gitlab.FeatureDefinition {
	result, e := c.FeatureDefinitions(x)
	errors.PanicOnError(e)

	return result
}
