package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) ProjectMergeRequests(
	project int64,
	all bool,
) ([]*merge_request.Request, error) {
	var result []*gitlab.BasicMergeRequest
	number := int64(1)

	for {
		o := &gitlab.ListProjectMergeRequestsOptions{
			State: new(constant.OpenedState),
			ListOptions: gitlab.ListOptions{
				PerPage: constant.PerPage100,
				Page:    number,
			},
		}

		if all {
			o.State = new("all")
		}

		page, _, e := c.client.MergeRequests.ListProjectMergeRequests(
			project,
			o,
		)

		if e != nil {
			return nil, e
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return merge_request.NewSlice(result), nil
}
