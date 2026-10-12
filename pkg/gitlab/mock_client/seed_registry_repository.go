package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) SeedRegistryRepository(
	project int64,
	identifier int64,
	path string,
) {
	c.registries = append(
		c.registries,
		registry_repository.New(
			&gitlab.RegistryRepository{
				ID:        identifier,
				ProjectID: project,
				Path:      path,
			},
		),
	)
}
