package mock_client

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeletePipeline(
	project int64,
	identifier int64,
) {
	errors.PanicOnError(c.DeletePipeline(project, identifier))
}
