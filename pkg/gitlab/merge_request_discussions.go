package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/discussion"

func (c *Client) MergeRequestDiscussions(
	project int64,
	identifier int64,
) ([]*discussion.Discussion, error) {
	result, _, e := c.client.Discussions.ListMergeRequestDiscussions(
		project,
		identifier,
		nil,
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return discussion.NewSlice(result), nil
}
