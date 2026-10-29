package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/task"

func (c *Client) AddTask(t *task.Task) {
	c.tasks = append(c.tasks, t)
}
