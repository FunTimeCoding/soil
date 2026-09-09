package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/diff"

func (c *Client) MergeRequestDiffs(
	_ int64,
	_ int64,
) ([]*diff.Diff, error) {
	return nil, nil
}
