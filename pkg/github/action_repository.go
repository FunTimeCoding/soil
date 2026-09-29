package github

import "github.com/funtimecoding/soil/pkg/github/repository"

func (c *Client) ActionRepository() ([]*repository.Repository, error) {
	var result []*repository.Repository
	u, e := c.User()

	if e != nil {
		return nil, e
	}

	codes, f := c.SearchCode(
		"actions/checkout user:%s in:file language:yaml",
		u.Name,
	)

	if f != nil {
		return nil, f
	}

	for _, o := range codes {
		result = append(result, repository.New(o.Raw.Repository))
	}

	return result, nil
}
