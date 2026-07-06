package gitlab

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustTrace(
	project int64,
	job int64,
) string {
	result, e := c.Trace(project, job)
	errors.PanicOnError(e)

	return result
}
