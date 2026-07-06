package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) MustAssignedMergeRequests(all bool) []*merge_request.Request {
	result, e := c.AssignedMergeRequests(all)
	errors.PanicOnError(e)

	return result
}
