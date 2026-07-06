package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustRegistryRepositories(
	project int64,
	panicOnForbidden bool,
) []*gitlab.RegistryRepository {
	result, e := c.RegistryRepositories(project, panicOnForbidden)
	errors.PanicOnError(e)

	return result
}
