package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustTags(project int64) []*gitlab.Tag {
	result, e := c.Tags(project)
	errors.PanicOnError(e)

	return result
}
