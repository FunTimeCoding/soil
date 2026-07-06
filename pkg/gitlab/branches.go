package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Branches(project int64) ([]*branch.Branch, error) {
	result, _, e := c.client.Branches.ListBranches(
		project,
		&gitlab.ListBranchesOptions{},
	)

	if e != nil {
		return nil, e
	}

	return branch.NewSlice(result), nil
}
