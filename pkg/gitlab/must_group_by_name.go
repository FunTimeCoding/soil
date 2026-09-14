package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustGroupByName(s string) *gitlab.Group {
	result, e := c.GroupByName(s)
	errors.PanicOnError(e)

	return result
}
