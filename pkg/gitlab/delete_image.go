package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) DeleteImage(
	x context.Context,
	project int64,
	repository int64,
	tag string,
) error {
	_, e := c.client.ContainerRegistry.DeleteRegistryRepositoryTag(
		project,
		repository,
		tag,
		gitlab.WithContext(x),
	)

	return wrapError(e)
}
