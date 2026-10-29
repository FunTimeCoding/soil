package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

func (c *Client) GetTasks(
	_ context.Context,
	taskType string,
) ([]*client.Task, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if taskType == "" {
		return c.TaskList, nil
	}

	var result []*client.Task

	for _, t := range c.TaskList {
		if t.Type == taskType {
			result = append(result, t)
		}
	}

	return result, nil
}
