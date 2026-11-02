package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustPackages(
	x context.Context,
	project int64,
	panicOnForbidden bool,
) []*gitlab.Package {
	result, e := c.Packages(x, project, panicOnForbidden)
	errors.PanicOnError(e)

	return result
}
