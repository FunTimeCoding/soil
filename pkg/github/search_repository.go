package github

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github/repository"
)

func (c *Client) SearchRepository(
	query string,
	a ...any,
) ([]*repository.Repository, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	result, _, e := c.client.Search.Repositories(c.context, query, nil)

	if e != nil {
		return nil, e
	}

	return repository.NewSlice(result.Repositories), nil
}
