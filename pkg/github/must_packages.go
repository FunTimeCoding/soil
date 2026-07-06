package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/container"
)

func (c *Client) MustPackages(user string) []*container.Container {
	result, e := c.Packages(user)
	errors.PanicOnError(e)

	return result
}
