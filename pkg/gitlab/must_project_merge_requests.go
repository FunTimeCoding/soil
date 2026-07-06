package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) MustProjectMergeRequests(
	project int64,
	all bool,
) []*merge_request.Request {
	result, e := c.ProjectMergeRequests(project, all)
	errors.PanicOnError(e)

	return result
}
