package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustFeatureDefinitions() []*gitlab.FeatureDefinition {
	result, e := c.FeatureDefinitions()
	errors.PanicOnError(e)

	return result
}
