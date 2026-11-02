package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustPackageFiles(
	x context.Context,
	project int64,
	packageIdentifier int64,
) []*gitlab.PackageFile {
	result, e := c.PackageFiles(x, project, packageIdentifier)
	errors.PanicOnError(e)

	return result
}
