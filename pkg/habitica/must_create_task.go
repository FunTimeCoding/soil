package habitica

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/habitica/request"
	"github.com/funtimecoding/soil/pkg/habitica/task"
)

func (c *Client) MustCreateTask(b *request.CreateTaskBody) *task.Task {
	result, e := c.CreateTask(b)
	errors.PanicOnError(e)

	return result
}
