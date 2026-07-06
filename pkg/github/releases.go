package github

import "github.com/funtimecoding/soil/pkg/github/release"

func (c *Client) Releases(
	namespace string,
	repository string,
) ([]*release.Release, error) {
	result, _, e := c.client.Repositories.ListReleases(
		c.context,
		namespace,
		repository,
		nil,
	)

	if e != nil {
		return nil, e
	}

	return release.NewSlice(result), nil
}
