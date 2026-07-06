package github

import "github.com/funtimecoding/soil/pkg/github/repository"

func (c *Client) ActionRepository() []*repository.Repository {
	var result []*repository.Repository

	for _, o := range c.MustSearchCode(
		"actions/checkout user:%s in:file language:yaml",
		c.MustUser().Name,
	) {
		result = append(result, repository.New(o.Raw.Repository))
	}

	return result
}
