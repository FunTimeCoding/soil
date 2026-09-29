package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) MustProjectsMergeRequests() []*merge_request.Request {
	result, e := c.ProjectsMergeRequests()
	errors.PanicOnError(e)

	return result
}
