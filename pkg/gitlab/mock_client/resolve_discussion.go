package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
)

func (c *Client) ResolveDiscussion(
	_ context.Context,
	_ int64,
	_ int64,
	thread string,
) error {
	if e := discussion.Check(thread); e != nil {
		return e
	}

	if c.failure != nil {
		return c.failure
	}

	c.resolved[thread] = true

	return nil
}
