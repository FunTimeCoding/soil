package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
)

func (c *Client) MustRegistryRepositories(
	x context.Context,
	project int64,
	panicOnForbidden bool,
) []*registry_repository.Repository {
	result, e := c.RegistryRepositories(x, project, panicOnForbidden)
	errors.PanicOnError(e)

	return result
}
