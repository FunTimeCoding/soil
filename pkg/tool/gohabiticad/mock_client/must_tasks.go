package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/habitica/task"
)

func (c *Client) MustTasks(taskType string) []*task.Task {
	result, e := c.Tasks(taskType)
	errors.PanicOnError(e)

	return result
}
