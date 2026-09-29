package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/repository"
)

func (c *Client) MustActionRepository() []*repository.Repository {
	result, e := c.ActionRepository()
	errors.PanicOnError(e)

	return result
}
