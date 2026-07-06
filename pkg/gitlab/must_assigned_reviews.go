package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) MustAssignedReviews(all bool) []*merge_request.Request {
	result, e := c.AssignedReviews(all)
	errors.PanicOnError(e)

	return result
}
