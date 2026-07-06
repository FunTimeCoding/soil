package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/task"
)

func (c *Client) MustTasks() []*task.Task {
	result, e := c.Tasks()
	errors.PanicOnError(e)

	return result
}
