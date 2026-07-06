package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/google/go-github/v88/github"
)

func (c *Client) MustRepositories(owner string) []*github.Repository {
	result, e := c.Repositories(owner)
	errors.PanicOnError(e)

	return result
}
