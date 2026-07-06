package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
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
		return nil, fmt.Errorf(
			"project not found: %d: %w",
			project,
			constant.ErrorNotFound,
		)
	}

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(
			"no merge request for branch %s in project %d: %w",
			branch,
			project,
			constant.ErrorNotFound,
		)
	}

	return merge_request.New(result[0]), nil
}
