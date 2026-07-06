package github

import (
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/image"
)

func (c *Client) PackageVersions(packageName string) ([]*image.Image, error) {
	result, _, e := c.client.Users.ListPackageVersions(
		c.context,
		constant.ContainerPackageType,
		packageName,
		nil,
	)

	if e != nil {
		return nil, e
	}

	return image.NewSlice(result), nil
}
