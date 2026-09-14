package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustPackages(
	project int64,
	panicOnForbidden bool,
) []*gitlab.Package {
	result, e := c.Packages(project, panicOnForbidden)
	errors.PanicOnError(e)

	return result
}
