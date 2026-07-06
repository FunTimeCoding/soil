package github

import "github.com/funtimecoding/soil/pkg/github/constant"

func (c *Client) DeletePackageVersion(
	packageName string,
	versionIdentifier int64,
) error {
	_, e := c.client.Users.PackageDeleteVersion(
		c.context,
		"",
		constant.ContainerPackageType,
		packageName,
		versionIdentifier,
	)

	return e
}
