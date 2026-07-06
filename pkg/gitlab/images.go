package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Images(
	project int64,
	repository int64,
) ([]*gitlab.RegistryRepositoryTag, error) {
	result, _, e := c.client.ContainerRegistry.ListRegistryRepositoryTags(
		project,
		repository,
		&gitlab.ListRegistryRepositoryTagsOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: constant.PerPage1000,
			},
		},
	)

	if e != nil {
		return nil, e
	}

	return image.Sort(result), nil
}
