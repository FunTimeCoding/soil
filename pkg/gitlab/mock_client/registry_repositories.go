package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
)

func (c *Client) RegistryRepositories(
	_ context.Context,
	project int64,
	_ bool,
) ([]*registry_repository.Repository, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	var result []*registry_repository.Repository

	for _, r := range c.registries {
		if r.ProjectIdentifier == project {
			result = append(result, r)
		}
	}

	return result, nil
}
