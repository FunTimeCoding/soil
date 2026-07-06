package habitica

import (
	"github.com/funtimecoding/soil/pkg/habitica/request"
	"github.com/funtimecoding/soil/pkg/habitica/task"
)

func (c *Client) CreateTask(b *request.CreateTaskBody) (*task.Task, error) {
	var result *task.Task

	return result, c.post("/tasks/user", b, &result)
}
