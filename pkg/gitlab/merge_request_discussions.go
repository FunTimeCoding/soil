package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MergeRequestDiscussions(
	x context.Context,
	project int64,
	identifier int64,
) ([]*discussion.Discussion, error) {
	result, _, e := c.client.Discussions.ListMergeRequestDiscussions(
		project,
		identifier,
		nil,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return discussion.NewSlice(result), nil
}
