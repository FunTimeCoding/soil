package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustCreateTag(
	project int64,
	name string,
	reference string,
	message string,
) *gitlab.Tag {
	result, e := c.CreateTag(project, name, reference, message)
	errors.PanicOnError(e)

	return result
}
