package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustPackageFiles(
	project int64,
	packageIdentifier int64,
) []*gitlab.PackageFile {
	result, e := c.PackageFiles(project, packageIdentifier)
	errors.PanicOnError(e)

	return result
}
