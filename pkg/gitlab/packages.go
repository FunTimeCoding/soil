package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Packages(
	project int64,
	panicOnForbidden bool,
) ([]*gitlab.Package, error) {
	result, r, e := c.client.Packages.ListProjectPackages(
		project,
		&gitlab.ListProjectPackagesOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
	)

	if !panicOnForbidden && r != nil && r.StatusCode == 403 {
		// Given correct token scope, this might be due to the GitLab server being configured wrong: https://forum.gitlab.com/t/cant-login-to-registry-due-to-denied-access-forbidden/63965/6
		errors.Warning("packages 403")

		return result, nil
	}

	return result, e
}
