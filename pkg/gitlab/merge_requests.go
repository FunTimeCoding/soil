package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_filter"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MergeRequests(
	x context.Context,
	f *merge_request_filter.Filter,
) ([]*merge_request.Request, error) {
	first := int64(f.Offset/f.Limit + 1)
	skip := f.Offset % f.Limit
	var result []*gitlab.BasicMergeRequest

	for page := first; page <= first+1; page++ {
		v, _, e := c.client.MergeRequests.ListProjectMergeRequests(
			f.Project,
			f.Options(page),
			gitlab.WithContext(x),
		)

		if e != nil {
			return nil, wrapError(e)
		}

		result = append(result, v...)

		if skip == 0 || len(v) < f.Limit {
			break
		}
	}

	if skip >= len(result) {
		return nil, nil
	}

	return merge_request.NewSlice(result[skip:min(skip+f.Limit, len(result))]), nil
}
