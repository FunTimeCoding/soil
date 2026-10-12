package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
	"slices"
)

func (c *Client) DeleteRegistryRepository(
	project int64,
	repository int64,
) error {
	c.registries = slices.DeleteFunc(
		c.registries,
		func(r *registry_repository.Repository) bool {
			return r.ProjectIdentifier == project && r.Identifier == repository
		},
	)

	return nil
}
