package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) RegistryRepositories(
	project int64,
	panicOnForbidden bool,
) ([]*gitlab.RegistryRepository, error) {
	var result []*gitlab.RegistryRepository
	number := int64(1)

	for {
		page, r, e := c.client.ContainerRegistry.ListProjectRegistryRepositories(
			project,
			&gitlab.ListProjectRegistryRepositoriesOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: constant.PerPage100,
					Page:    number,
				},
			},
		)

		if !panicOnForbidden && r != nil && r.StatusCode == 403 {
			// Given correct token scope, this might be due to the GitLab server being configured wrong: https://forum.gitlab.com/t/cant-login-to-registry-due-to-denied-access-forbidden/63965/6
			errors.Warning("registry repositories 403")

			return result, nil
		}

		if e != nil {
			return nil, e
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return result, nil
}
