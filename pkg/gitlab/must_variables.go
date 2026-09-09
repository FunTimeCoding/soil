package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
)

func (c *Client) MustVariables(project int64) []*variable.Variable {
	result, e := c.Variables(project)
	errors.PanicOnError(e)

	return result
}
