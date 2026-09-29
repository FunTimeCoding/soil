package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/response"
)

func (c *Client) MustGraphRunner(identifier int64) *response.Runner {
	result, e := c.GraphRunner(identifier)
	errors.PanicOnError(e)

	return result
}
