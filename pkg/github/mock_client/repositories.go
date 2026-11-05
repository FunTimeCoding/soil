package mock_client

import (
	"github.com/google/go-github/v93/github"
	"maps"
	"slices"
)

func (c *Client) Repositories(_ string) ([]*github.Repository, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []*github.Repository

	for _, name := range slices.Sorted(maps.Keys(c.runs)) {
		result = append(result, &github.Repository{Name: new(name)})
	}

	return result, nil
}
