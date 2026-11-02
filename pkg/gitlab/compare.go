package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Compare(
	x context.Context,
	project int64,
	from, to string,
) (*gitlab.Compare, error) {
	result, _, e := c.client.Repositories.Compare(
		project,
		&gitlab.CompareOptions{From: new(from), To: new(to)},
		gitlab.WithContext(x),
	)

	return result, wrapError(e)
}
