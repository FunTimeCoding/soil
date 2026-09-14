package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustFeatures() []*gitlab.Feature {
	result, e := c.Features()
	errors.PanicOnError(e)

	return result
}
