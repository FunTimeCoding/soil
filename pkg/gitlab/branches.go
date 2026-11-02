package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Branches(
	x context.Context,
	project int64,
) ([]*branch.Branch, error) {
	result, _, e := c.client.Branches.ListBranches(
		project,
		&gitlab.ListBranchesOptions{},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return branch.NewSlice(result), nil
}
