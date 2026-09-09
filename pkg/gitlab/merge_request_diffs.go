package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/diff"

func (c *Client) MergeRequestDiffs(
	project int64,
	identifier int64,
) ([]*diff.Diff, error) {
	result, _, e := c.client.MergeRequests.ListMergeRequestDiffs(
		project,
		identifier,
		nil,
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return diff.NewMergeRequestSlice(result), nil
}
