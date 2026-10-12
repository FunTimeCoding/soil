package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/registry_repository"

func (c *Client) RegistryRepositories(
	project int64,
	_ bool,
) ([]*registry_repository.Repository, error) {
	var result []*registry_repository.Repository

	for _, r := range c.registries {
		if r.ProjectIdentifier == project {
			result = append(result, r)
		}
	}

	return result, nil
}
