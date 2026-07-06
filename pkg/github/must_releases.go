package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/release"
)

func (c *Client) MustReleases(
	namespace string,
	repository string,
) []*release.Release {
	result, e := c.Releases(namespace, repository)
	errors.PanicOnError(e)

	return result
}
