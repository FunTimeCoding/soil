package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) PackageFiles(
	x context.Context,
	project int64,
	packageIdentifier int64,
) ([]*gitlab.PackageFile, error) {
	result, _, e := c.client.Packages.ListPackageFiles(
		project,
		packageIdentifier,
		nil,
		gitlab.WithContext(x),
	)

	return result, wrapError(e)
}
