package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/diff"

func (c *Client) CommitDiff(
	_ int64,
	_ string,
	_ int64,
) ([]*diff.Diff, error) {
	return nil, nil
}
