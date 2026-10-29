package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/task"

func (c *Client) Tasks(taskType string) ([]*task.Task, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if taskType == "" {
		return c.tasks, nil
	}

	var result []*task.Task

	for _, t := range c.tasks {
		if t.Type == taskType {
			result = append(result, t)
		}
	}

	return result, nil
}
