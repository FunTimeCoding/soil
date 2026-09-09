package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) ListCommits(
	project int64,
	reference string,
	limit int64,
) ([]*commit.Commit, error) {
	if limit == 0 {
		limit = 20
	}

	o := &gitlab.ListCommitsOptions{
		ListOptions: gitlab.ListOptions{PerPage: limit},
	}

	if reference != "" {
		o.RefName = &reference
	}

	result, _, e := c.client.Commits.ListCommits(project, o)

	if e != nil {
		return nil, wrapError(e)
	}

	return commit.NewSlice(result), nil
}
