package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) ProjectMergeRequests(
	project int64,
	state string,
) ([]*merge_request.Request, error) {
	if state == "" {
		state = constant.OpenedState
	}

	var result []*gitlab.BasicMergeRequest
	number := int64(1)

	for {
		o := &gitlab.ListProjectMergeRequestsOptions{
			State: &state,
			ListOptions: gitlab.ListOptions{
				PerPage: constant.PerPage100,
				Page:    number,
			},
		}
		page, _, e := c.client.MergeRequests.ListProjectMergeRequests(
			project,
			o,
		)

		if e != nil {
			return nil, wrapError(e)
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return merge_request.NewSlice(result), nil
}
