package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) CreateTag(
	x context.Context,
	project int64,
	name string,
	reference string,
	message string,
) (*tag.Tag, error) {
	result, _, e := c.client.Tags.CreateTag(
		project,
		&gitlab.CreateTagOptions{
			TagName: new(name),
			Ref:     new(reference),
			Message: new(message),
		},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return tag.New(result), nil
}
