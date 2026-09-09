package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) CommitDiff(
	project int64,
	sha string,
	limit int64,
) ([]*diff.Diff, error) {
	if limit == 0 {
		limit = 100
	}

	result, _, e := c.client.Commits.GetCommitDiff(
		project,
		sha,
		&gitlab.GetCommitDiffOptions{
			ListOptions: gitlab.ListOptions{PerPage: limit},
		},
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return diff.NewSlice(result), nil
}
