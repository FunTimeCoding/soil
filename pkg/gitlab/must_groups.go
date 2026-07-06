package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustGroups() []*gitlab.Group {
	result, e := c.Groups()
	errors.PanicOnError(e)

	return result
}
