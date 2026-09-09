package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/discussion"

func (c *Client) MergeRequestDiscussions(
	_ int64,
	_ int64,
) ([]*discussion.Discussion, error) {
	return nil, nil
}
