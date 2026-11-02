package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Tags(
	x context.Context,
	project int64,
) ([]*tag.Tag, error) {
	result, _, e := c.client.Tags.ListTags(
		project,
		&gitlab.ListTagsOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return tag.NewSlice(result), nil
}
