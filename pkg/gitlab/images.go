package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Images(
	x context.Context,
	project int64,
	repository int64,
) ([]*image.Image, error) {
	result, _, e := c.client.ContainerRegistry.ListRegistryRepositoryTags(
		project,
		repository,
		&gitlab.ListRegistryRepositoryTagsOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return image.Sort(image.NewSlice(result)), nil
}
