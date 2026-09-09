package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Tags(project int64) ([]*tag.Tag, error) {
	result, _, e := c.client.Tags.ListTags(
		project,
		&gitlab.ListTagsOptions{
			ListOptions: gitlab.ListOptions{PerPage: constant.PerPage1000},
		},
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return tag.NewSlice(result), nil
}
