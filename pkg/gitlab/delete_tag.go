package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) DeleteTag(
	x context.Context,
	project int64,
	name string,
) error {
	_, e := c.client.Tags.DeleteTag(project, name, gitlab.WithContext(x))

	return wrapError(e)
}
