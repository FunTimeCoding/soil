package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Groups(x context.Context) ([]*gitlab.Group, error) {
	result, _, e := c.client.Groups.ListGroups(nil, gitlab.WithContext(x))

	return result, wrapError(e)
}
