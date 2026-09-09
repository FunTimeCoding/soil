package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
)

func (c *Client) MustRegistryRepositories(
	project int64,
	panicOnForbidden bool,
) []*registry_repository.Repository {
	result, e := c.RegistryRepositories(project, panicOnForbidden)
	errors.PanicOnError(e)

	return result
}
