package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
)

func (c *Client) MustBranches(project int64) []*branch.Branch {
	result, e := c.Branches(project)
	errors.PanicOnError(e)

	return result
}
