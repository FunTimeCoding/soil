package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) PackageFiles(
	project int64,
	packageIdentifier int64,
) ([]*gitlab.PackageFile, error) {
	result, _, e := c.client.Packages.ListPackageFiles(
		project,
		packageIdentifier,
		nil,
	)

	return result, wrapError(e)
}
