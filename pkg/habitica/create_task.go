package habitica

import (
	"github.com/funtimecoding/soil/pkg/habitica/request"
	"github.com/funtimecoding/soil/pkg/habitica/task"
)

func (c *Client) CreateTask(b *request.CreateTaskBody) (*task.Task, error) {
	var result *task.Task
	e := c.basic.Post("/tasks/user", nil, b, &result)

	return result, e
}
