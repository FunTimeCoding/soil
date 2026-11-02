package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) DeletePackage(
	x context.Context,
	project int64,
	repository int64,
) error {
	_, e := c.client.Packages.DeleteProjectPackage(
		project,
		repository,
		gitlab.WithContext(x),
	)

	return wrapError(e)
}
