package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
)

func (c *Client) MergeRequestDiscussions(
	_ context.Context,
	_ int64,
	_ int64,
) ([]*discussion.Discussion, error) {
	return nil, nil
}
