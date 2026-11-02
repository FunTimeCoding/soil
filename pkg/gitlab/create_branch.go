package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) CreateBranch(
	x context.Context,
	project int64,
	name string,
	reference string,
) (*branch.Branch, error) {
	result, _, e := c.client.Branches.CreateBranch(
		project,
		&gitlab.CreateBranchOptions{Branch: &name, Ref: &reference},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return branch.New(result), nil
}
