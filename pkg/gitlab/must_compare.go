package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustCompare(
	project int64,
	from, to string,
) *gitlab.Compare {
	result, e := c.Compare(project, from, to)
	errors.PanicOnError(e)

	return result
}
