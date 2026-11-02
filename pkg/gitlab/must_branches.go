package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
)

func (c *Client) MustBranches(
	x context.Context,
	project int64,
) []*branch.Branch {
	result, e := c.Branches(x, project)
	errors.PanicOnError(e)

	return result
}
