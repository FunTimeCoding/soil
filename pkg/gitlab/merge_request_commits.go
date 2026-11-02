package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MergeRequestCommits(
	x context.Context,
	project int64,
	request int64,
) ([]*commit.Commit, error) {
	result, _, e := c.client.MergeRequests.GetMergeRequestCommits(
		project,
		request,
		&gitlab.GetMergeRequestCommitsOptions{},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return commit.NewSlice(result), nil
}
