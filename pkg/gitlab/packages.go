package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Packages(
	x context.Context,
	project int64,
	panicOnForbidden bool,
) ([]*gitlab.Package, error) {
	result, r, e := c.client.Packages.ListProjectPackages(
		project,
		&gitlab.ListProjectPackagesOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
		gitlab.WithContext(x),
	)

	if !panicOnForbidden && r != nil && r.StatusCode == 403 {
		errors.Warning("packages 403")

		return result, nil
	}

	return result, wrapError(e)
}
