package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) DeleteProjectVariable(
	x context.Context,
	project int64,
	key string,
) error {
	_, e := c.client.ProjectVariables.RemoveVariable(
		project,
		key,
		nil,
		gitlab.WithContext(x),
	)

	return wrapError(e)
}
