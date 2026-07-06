package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustVariables(project int64) []*gitlab.ProjectVariable {
	result, e := c.Variables(project)
	errors.PanicOnError(e)

	return result
}
