package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/repository"
)

func (c *Client) MustSearchRepository(
	query string,
	a ...any,
) []*repository.Repository {
	result, e := c.SearchRepository(query, a...)
	errors.PanicOnError(e)

	return result
}
