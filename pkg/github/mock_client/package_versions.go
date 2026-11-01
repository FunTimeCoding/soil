package mock_client

import "github.com/funtimecoding/soil/pkg/github/image"

func (c *Client) PackageVersions(_ string) ([]*image.Image, error) {
	return nil, c.Failure
}
