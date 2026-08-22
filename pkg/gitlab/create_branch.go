package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CreateBranch(
	project int64,
	name string,
	reference string,
) (*gitlab.Branch, error) {
	result, _, e := c.client.Branches.CreateBranch(
		project,
		&gitlab.CreateBranchOptions{Branch: &name, Ref: &reference},
	)

	return result, wrapError(e)
}
