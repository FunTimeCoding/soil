package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) CreateBranch(
	project int64,
	name string,
	reference string,
) (*branch.Branch, error) {
	result, _, e := c.client.Branches.CreateBranch(
		project,
		&gitlab.CreateBranchOptions{Branch: &name, Ref: &reference},
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return branch.New(result), nil
}
