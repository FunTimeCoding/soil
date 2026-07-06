package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Jobs() []*gojenkins.Job {
	if false {
		type contextKey string
		c.context = context.WithValue(
			c.context,
			contextKey("debug"),
			true,
		)
	}

	result, e := c.client.GetAllJobs(c.context)
	errors.PanicOnError(e)

	return result
}
