package github

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeletePackageVersion(
	packageName string,
	versionIdentifier int64,
) {
	errors.PanicOnError(
		c.DeletePackageVersion(
			packageName,
			versionIdentifier,
		),
	)
}
