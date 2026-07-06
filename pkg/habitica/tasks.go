package habitica

import (
	"github.com/funtimecoding/soil/pkg/habitica/task"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) Tasks(taskType string) ([]*task.Task, error) {
	path := "/tasks/user"

	if taskType != "" {
		path = join.Empty(path, "?type=", taskType)
	}

	var result []*task.Task

	return result, c.get(path, &result)
}
