package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MergeRequestCommits(
	project int64,
	request int64,
) ([]*commit.Commit, error) {
	result, _, e := c.client.MergeRequests.GetMergeRequestCommits(
		project,
		request,
		&gitlab.GetMergeRequestCommitsOptions{},
	)

	if e != nil {
		return nil, e
	}

	return commit.NewSlice(result), nil
}
