package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MergeRequestDiffs(
	x context.Context,
	project int64,
	identifier int64,
) ([]*diff.Diff, error) {
	result, _, e := c.client.MergeRequests.ListMergeRequestDiffs(
		project,
		identifier,
		nil,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return diff.NewMergeRequestSlice(result), nil
}
