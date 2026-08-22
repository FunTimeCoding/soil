package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Tree(project int64) ([]*gitlab.TreeNode, error) {
	result, r, e := c.client.Repositories.ListTree(
		project,
		&gitlab.ListTreeOptions{},
	)

	if r != nil && r.StatusCode == 404 {
		// Do not panic
		return []*gitlab.TreeNode{}, nil
	}

	return result, wrapError(e)
}
