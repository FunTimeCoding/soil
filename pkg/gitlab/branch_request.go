package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) BranchRequest(
	project int64,
	branch string,
) (*merge_request.Request, error) {
	result, r, e := c.client.MergeRequests.ListProjectMergeRequests(
		project,
		&gitlab.ListProjectMergeRequestsOptions{
			SourceBranch: new(branch),
			ListOptions:  gitlab.ListOptions{PerPage: 1},
		},
	)

	if r != nil && r.StatusCode == 404 {
		return nil, not_found.New("project", project)
	}

	if e != nil {
		return nil, wrapError(e)
	}

	if len(result) == 0 {
		return nil, not_found.Format(
			"no merge request for branch %s in project %d",
			branch,
			project,
		)
	}

	return merge_request.New(result[0]), nil
}
