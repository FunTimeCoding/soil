package gitlab

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeletePipeline(
	project int64,
	pipeline int64,
) {
	errors.PanicOnError(c.DeletePipeline(project, pipeline))
}
