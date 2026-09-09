package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/registry_repository"

func (c *Client) RegistryRepositories(
	_ int64,
	_ bool,
) ([]*registry_repository.Repository, error) {
	return nil, nil
}
