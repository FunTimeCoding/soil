package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) ResolveDiscussion(
	x context.Context,
	project int64,
	identifier int64,
	thread string,
) error {
	if e := discussion.Check(thread); e != nil {
		return e
	}

	_, _, e := c.client.Discussions.ResolveMergeRequestDiscussion(
		project,
		identifier,
		thread,
		&gitlab.ResolveMergeRequestDiscussionOptions{Resolved: new(true)},
		gitlab.WithContext(x),
	)

	if e != nil {
		return wrapError(e)
	}

	return nil
}
